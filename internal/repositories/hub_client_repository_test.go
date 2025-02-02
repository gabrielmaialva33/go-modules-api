package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-modules-api/internal/models"
	"go-modules-api/internal/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestHubClientRepository_Pagination(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewHubClientRepository(gormDB)

	testCases := []struct {
		name         string
		search       string
		active       *bool
		sortField    string
		sortOrder    string
		page         int
		pageSize     int
		rowsReturned int
		totalCount   int64
		errExpected  bool
	}{
		{"valid_no_filter", "", nil, "name", "asc", 1, 2, 2, 10, false},
		{"valid_with_search", "demo", nil, "name", "asc", 1, 1, 1, 1, false},
		{"invalid_sort_order", "", nil, "name", "invalid", 1, 2, 0, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			countRegex := `SELECT count\(\*\) FROM "hub_clients"(?: WHERE .+)?`
			mock.ExpectQuery(countRegex).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(tc.totalCount))

			findRegex := `SELECT \* FROM "hub_clients"(?: WHERE .+)?`
			if tc.rowsReturned > 0 {
				rows := sqlmock.NewRows([]string{"id"})
				for i := 0; i < tc.rowsReturned; i++ {
					rows.AddRow(i + 1)
				}
				mock.ExpectQuery(findRegex).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(findRegex).WillReturnRows(sqlmock.NewRows([]string{"id"}))
			}

			clients, total, err := repo.Pagination(tc.search, tc.active, tc.sortField, tc.sortOrder, tc.page, tc.pageSize)
			if tc.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, clients, tc.rowsReturned)
				assert.Equal(t, tc.totalCount, total)
			}
		})
	}
}

func TestHubClientRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewHubClientRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		search      string
		active      *bool
		sortField   string
		sortOrder   string
		expectError bool
	}{
		{"success", func() {
			mock.ExpectQuery(`SELECT \* FROM "hub_clients"`).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		}, "", nil, "name", "asc", false},
		{"db_error", func() {
			mock.ExpectQuery(`SELECT \* FROM "hub_clients"`).
				WillReturnError(gorm.ErrInvalidData)
		}, "", nil, "name", "asc", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareMock()
			clients, err := repo.GetAll(tc.search, tc.active, tc.sortField, tc.sortOrder)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, clients, 1)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHubClientRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewHubClientRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		id          uint
		expectError bool
	}{
		{"success", func() {
			// Atualize o regex para incluir a condição is_deleted = false e o LIMIT $3.
			mock.ExpectQuery(`SELECT \* FROM "hub_clients" WHERE "hub_clients"."id" = \$1 AND is_deleted = \$2 ORDER BY "hub_clients"."id" LIMIT \$3`).
				WithArgs(1, false, 1).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		}, 1, false},
		{"not_found", func() {
			mock.ExpectQuery(`SELECT \* FROM "hub_clients" WHERE "hub_clients"."id" = \$1 AND is_deleted = \$2 ORDER BY "hub_clients"."id" LIMIT \$3`).
				WithArgs(1, false, 1).
				WillReturnRows(sqlmock.NewRows([]string{"id"}))
		}, 1, true},
		{"db_error", func() {
			mock.ExpectQuery(`SELECT \* FROM "hub_clients" WHERE "hub_clients"."id" = \$1 AND is_deleted = \$2 ORDER BY "hub_clients"."id" LIMIT \$3`).
				WithArgs(1, false, 1).
				WillReturnError(gorm.ErrInvalidData)
		}, 1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareMock()
			_, err := repo.GetByID(tc.id)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHubClientRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewHubClientRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		input       *models.HubClient
		expectError bool
	}{
		{"success", func() {
			mock.ExpectBegin()
			mock.ExpectQuery(`INSERT INTO "hub_clients"`).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()
		}, &models.HubClient{Name: "New Client"}, false},
		{"db_error", func() {
			mock.ExpectBegin()
			mock.ExpectQuery(`INSERT INTO "hub_clients"`).
				WillReturnError(gorm.ErrInvalidData)
			mock.ExpectRollback()
		}, &models.HubClient{Name: "Failing Client"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareMock()
			err := repo.Create(tc.input)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHubClientRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewHubClientRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		input       *models.HubClient
		expectError bool
	}{
		{"success", func() {
			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE "hub_clients" SET`).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}, &models.HubClient{BaseID: models.BaseID{ID: 1}, Name: "Updated Client"}, false},
		{"db_error", func() {
			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE "hub_clients" SET`).
				WillReturnError(gorm.ErrInvalidData)
			mock.ExpectRollback()
		}, &models.HubClient{BaseID: models.BaseID{ID: 1}, Name: "Fail Update"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareMock()
			err := repo.Update(tc.input)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHubClientRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Abra o GORM com o driver postgres
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewHubClientRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		id          uint
		expectError bool
	}{
		{
			"success",
			func() {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM "hub_clients" WHERE "hub_clients"."id" = \$1 AND is_deleted = \$2`).
					WithArgs(1, false).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			1, false,
		},
		{
			"db_error",
			func() {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM "hub_clients" WHERE "hub_clients"."id" = \$1 AND is_deleted = \$2`).
					WithArgs(1, false).
					WillReturnError(gorm.ErrInvalidTransaction)
				mock.ExpectRollback()
			},
			1, true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.prepareMock()
			err := repo.Delete(tc.id)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
