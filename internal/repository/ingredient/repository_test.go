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

const testDbName = "test.sqlite3"

var defaultIngredients = []*model.Ingredient{
	{ID: 1, Name: "Ingredient 1"},
	{ID: 2, Name: "Ingredient 2"},
	{ID: 3, Name: "Ingredient 3"},
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", testDbName)
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

func DBSetup(db *sql.DB) error {
	q := `CREATE TABLE IF NOT EXISTS ingredients (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(64) NULL);`
	if _, err := db.Exec(q); err != nil {
		return fmt.Errorf("cannot create table: %w", err)
	}

	q = `INSERT INTO ingredients (name) VALUES(?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return fmt.Errorf("could not prepare insert query: %w", err)
	}

	defer func(stmt *sql.Stmt) { _ = stmt.Close() }(stmt)

	for _, i := range defaultIngredients {
		r, err := stmt.Exec(i.Name)
		if err != nil {
			return fmt.Errorf("cound not insert data: %w", err)
		}

		if num, err := r.RowsAffected(); err != nil || num != 1 {
			return fmt.Errorf("non-effective insert: %w", err)
		}
	}

	return nil
}

func DBCleanup(db *sql.DB) error {
	q := `DROP TABLE ingredients;`
	if _, err := db.Exec(q); err != nil {
		return fmt.Errorf("cannot drop the table: %w", err)
	}

	return nil
}

func TestMain(m *testing.M) {
	BeforeAll()

	code := m.Run()

	AfterAll()

	os.Exit(code)
}

func BeforeAll() {
	fmt.Println("BeforeAll")
}

func AfterAll() {
	_ = os.Remove(testDbName)
}

func BeforeEach(t *testing.T) {
	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	if err = DBSetup(db); err != nil {
		t.Fatal(err)
	}

	if err := closeDB(db); err != nil {
		t.Fatal(err)
	}
}

func AfterEach(t *testing.T) {
	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	if err := DBCleanup(db); err != nil {
		t.Fatal(err)
	}

	if err := closeDB(db); err != nil {
		t.Fatal(err)
	}

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
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should return all the ingredients in the database", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := ingredient.NewInDBIngredientRepository(db)

		ingredients, err := repo.GetAll()
		assert.NoError(err)

		assert.Equal(defaultIngredients, ingredients)
	})
}

func TestInDBIngredientRepo_Create(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should insert the given ingredient into the database", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := ingredient.NewInDBIngredientRepository(db)

		i := &model.Ingredient{ID: 4, Name: "Ingredient 4"}
		err = repo.Create(i)
		assert.NoError(err)

		ingredients, err := repo.GetAll()
		assert.NoError(err)

		assert.Contains(ingredients, i)
	})
}

func TestInDBIngredientRepo_Remove(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should remote the given ingredient from the database", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := ingredient.NewInDBIngredientRepository(db)

		err = repo.Remove(3)
		assert.NoError(err)

		ingredients, err := repo.GetAll()
		assert.NoError(err)

		expectedIngredients := defaultIngredients[0 : len(defaultIngredients)-1]
		assert.Equal(expectedIngredients, ingredients)
	})
}
