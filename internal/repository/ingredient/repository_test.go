package ingredient_test

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	assertpkg "github.com/stretchr/testify/assert"
	"os"
	"platter/internal/model"
	"platter/internal/repository/ingredient"
	"platter/internal/test"
	"testing"
)

var expectedIngredients = []*model.Ingredient{
	model.NewIngredient(1, "Ingredient 1"),
	model.NewIngredient(2, "Ingredient 2"),
	model.NewIngredient(3, "Ingredient 3"),
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "test.sqlite")
	if err != nil {
		return nil, fmt.Errorf("cannot open the database: %w", err)
	}

	return db, nil
}

func closeDB(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("cannot close the database: %w", err)
	}

	return nil
}

func DBSetup() (*sql.DB, error) {
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("cannot open the database: %w", err)
	}

	q := `CREATE TABLE IF NOT EXISTS ingredients (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(64) NULL);`
	if _, err = db.Exec(q); err != nil {
		return nil, fmt.Errorf("cannot create table: %w", err)
	}

	q = `INSERT INTO ingredients (name) VALUES(?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, fmt.Errorf("could not prepare insert query: %w", err)
	}

	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	for _, i := range expectedIngredients {
		r, err := stmt.Exec(i.Name)
		if err != nil {
			return nil, fmt.Errorf("cound not insert data: %w", err)
		}

		if num, err := r.RowsAffected(); err != nil || num != 1 {
			return nil, fmt.Errorf("non-effective insert: %w", err)
		}
	}

	return db, nil
}

func DBTeardown(db *sql.DB) error {
	if err := closeDB(db); err != nil {
		return fmt.Errorf("cannot close the database: %w", err)
	}

	_ = os.Remove("test.sqlite")

	return nil
}

func TestMain(m *testing.M) {
	// Setup
	db, err := DBSetup()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := m.Run()

	// Teardown
	if err := DBTeardown(db); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(code)
}

func TestInDBIngredientRepo_New(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should return a DB ingredient repository with no ingredients", func(t *testing.T) {
		repo := ingredient.NewInDBIngredientRepository(nil)
		assert.NotNil(repo)
	})
}

func TestInDBIngredientRepo_GetAll(t *testing.T) {
	assert := assertpkg.New(t)

	t.Run("should return all the ingredients in the database", func(t *testing.T) {
		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := ingredient.NewInDBIngredientRepository(db)

		ingredients, err := repo.GetAll()
		assert.NoError(err)

		assert.Equal(expectedIngredients, ingredients)
	})
}
