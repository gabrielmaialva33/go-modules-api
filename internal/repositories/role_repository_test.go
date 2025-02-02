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

func TestRoleRepository_Pagination(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewRoleRepository(gormDB)

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
		{"valid_with_search", "admin", nil, "name", "asc", 1, 1, 1, 1, false},
		{"invalid_sort_order", "", nil, "name", "invalid", 1, 2, 0, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			countRegex := `SELECT count\(\*\) FROM "roles"(?: WHERE .+)?`
			mock.ExpectQuery(countRegex).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(tc.totalCount))

			findRegex := `SELECT \* FROM "roles"(?: WHERE .+)?`
			if tc.rowsReturned > 0 {
				rows := sqlmock.NewRows([]string{"id"})
				for i := 0; i < tc.rowsReturned; i++ {
					rows.AddRow(i + 1)
				}
				mock.ExpectQuery(findRegex).WillReturnRows(rows)
			} else {
				mock.ExpectQuery(findRegex).WillReturnRows(sqlmock.NewRows([]string{"id"}))
			}

			roles, total, err := repo.Pagination(tc.search, tc.active, tc.sortField, tc.sortOrder, tc.page, tc.pageSize)
			if tc.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, roles, tc.rowsReturned)
				assert.Equal(t, tc.totalCount, total)
			}
		})
	}
}

func TestRoleRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewRoleRepository(gormDB)

	testCases := []struct {
		name        string
		search      string
		active      *bool
		sortField   string
		sortOrder   string
		roles       []models.Role
		expectError bool
	}{
		{"valid_no_filter", "", nil, "name", "asc", []models.Role{{BaseID: models.BaseID{ID: 1}, Name: "Admin"}}, false},
		{"valid_with_search", "admin", nil, "name", "asc", []models.Role{{BaseID: models.BaseID{ID: 1}, Name: "Admin"}}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id", "name"})
			// Supondo que o valor do ID seja 1
			for range tc.roles {
				rows.AddRow(1, "Admin")
			}
			mock.ExpectQuery(`SELECT \* FROM "roles"`).WillReturnRows(rows)

			roles, err := repo.GetAll(tc.search, tc.active, tc.sortField, tc.sortOrder)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.roles, roles)
			}
		})
	}
}

func TestRoleRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewRoleRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		id          uint
		expectError bool
	}{
		{"success", func() {
			mock.ExpectQuery(`SELECT \* FROM "roles" WHERE "roles"."id" = \$1 AND is_deleted = \$2 ORDER BY "roles"."id" LIMIT \$3`).
				WithArgs(1, false, 1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Admin"))
		}, 1, false},
		{"not_found", func() {
			mock.ExpectQuery(`SELECT \* FROM "roles" WHERE "roles"."id" = \$1 AND is_deleted = \$2 ORDER BY "roles"."id" LIMIT \$3`).
				WithArgs(1, false, 1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
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

func TestRoleRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewRoleRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		input       *models.Role
		expectError bool
	}{
		{"success", func() {
			mock.ExpectBegin()
			mock.ExpectQuery(`INSERT INTO "roles"`).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()
		}, &models.Role{Name: "New Role"}, false},
		{"db_error", func() {
			mock.ExpectBegin()
			mock.ExpectQuery(`INSERT INTO "roles"`).
				WillReturnError(gorm.ErrInvalidData)
			mock.ExpectRollback()
		}, &models.Role{Name: "Failing Role"}, true},
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

func TestRoleRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewRoleRepository(gormDB)

	testCases := []struct {
		name        string
		prepareMock func()
		input       *models.Role
		expectError bool
	}{
		{"success", func() {
			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE "roles" SET`).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}, &models.Role{BaseID: models.BaseID{ID: 1}, Name: "Updated Role"}, false},
		{"db_error", func() {
			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE "roles" SET`).
				WillReturnError(gorm.ErrInvalidData)
			mock.ExpectRollback()
		}, &models.Role{BaseID: models.BaseID{ID: 1}, Name: "Fail Update"}, true},
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

func TestRoleRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repositories.NewRoleRepository(gormDB)

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
				mock.ExpectExec(`DELETE FROM "roles" WHERE "roles"."id" = \$1 AND is_deleted = \$2`).
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
				mock.ExpectExec(`DELETE FROM "roles" WHERE "roles"."id" = \$1 AND is_deleted = \$2`).
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
