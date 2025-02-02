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

func TestPagination(t *testing.T) {
	// Cria uma conexão sqlmock e o respectivo *gorm.DB
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
		// Note que o case "invalid_sort_order" aqui não gera erro, pois o repositório
		// altera valores inválidos para "asc". Se desejar testar erro, ajuste o comportamento.
		{"valid_no_filter", "", nil, "name", "asc", 1, 2, 2, 10, false},
		{"valid_with_search", "demo", nil, "name", "asc", 1, 1, 1, 1, false},
		{"invalid_sort_order", "", nil, "name", "invalid", 1, 2, 0, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Expectativa para a query de contagem:
			// GORM gera uma query do tipo: SELECT count(*) FROM "hub_clients" [WHERE ...]
			countRegex := `SELECT count\(\*\) FROM "hub_clients"(?: WHERE .+)?`
			mock.ExpectQuery(countRegex).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(tc.totalCount))

			// Expectativa para a query que busca os registros:
			// GORM gera uma query do tipo: SELECT * FROM "hub_clients" [WHERE ...]
			findRegex := `SELECT \* FROM "hub_clients"(?: WHERE .+)?`
			if tc.rowsReturned > 0 {
				rows := sqlmock.NewRows([]string{"id"})
				for i := 0; i < tc.rowsReturned; i++ {
					rows.AddRow(i + 1)
				}
				mock.ExpectQuery(findRegex).WillReturnRows(rows)
			} else {
				// Se não há linhas, retornamos um conjunto vazio (ao invés de erro)
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

func TestCreate(t *testing.T) {
	// Cria a conexão sqlmock e o *gorm.DB
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
			// No PostgreSQL o GORM gera uma query com RETURNING "id",
			// que deve ser capturada com ExpectQuery e não ExpectExec.
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

func TestUpdate(t *testing.T) {
	// Cria a conexão sqlmock e o *gorm.DB
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
		}, &models.HubClient{ID: 1, Name: "Updated Client"}, false},
		{"db_error", func() {
			mock.ExpectBegin()
			mock.ExpectExec(`UPDATE "hub_clients" SET`).
				WillReturnError(gorm.ErrInvalidData)
			mock.ExpectRollback()
		}, &models.HubClient{ID: 1, Name: "Fail Update"}, true},
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

func TestDelete(t *testing.T) {
	// Cria a conexão sqlmock e o *gorm.DB
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
			mock.ExpectBegin()
			mock.ExpectExec(`DELETE FROM "hub_clients"`).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}, 1, false},
		{"db_error", func() {
			mock.ExpectBegin()
			mock.ExpectExec(`DELETE FROM "hub_clients"`).
				WillReturnError(gorm.ErrInvalidTransaction)
			mock.ExpectRollback()
		}, 1, true},
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
