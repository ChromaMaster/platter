package dish_test

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	assertpkg "github.com/stretchr/testify/assert"
	"os"
	"platter/internal/model"
	"platter/internal/repository/dish"
	"platter/internal/test"
	"testing"
)

const testDbName = "test.sqlite3"

var defaultDishes = []*model.Dish{
	model.NewDish(1, "Dish 1"),
	model.NewDish(2, "Dish 2"),
	model.NewDish(3, "Dish 3"),
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
	q := `CREATE TABLE IF NOT EXISTS dishes (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(64) NULL);`
	if _, err := db.Exec(q); err != nil {
		return fmt.Errorf("cannot create table: %w", err)
	}

	q = `INSERT INTO dishes (name) VALUES(?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return fmt.Errorf("could not prepare insert query: %w", err)
	}

	defer func(stmt *sql.Stmt) { _ = stmt.Close() }(stmt)

	for _, d := range defaultDishes {
		r, err := stmt.Exec(d.Name)
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
	q := `DROP TABLE dishes;`
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

func TestInDBDishRepo_New(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should return a DB dish repository", func(t *testing.T) {
		repo := dish.NewInDBRepository(nil)
		assert.NotNil(repo)
	})
}

func TestInDBDishRepository_Init(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should create the dish table in the database", func(t *testing.T) {
		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := dish.NewInDBRepository(db)
		err = repo.Init()
		assert.NoError(err)

		q := `SELECT name FROM sqlite_master WHERE type='table' AND name='dishes';`
		rows, err := db.Query(q)
		assert.NoError(err)

		count := 0
		for rows.Next() {
			count++
		}

		assert.Equal(1, count)
	})
}

func TestInDBDishRepo_GetAll(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should return all the dishes in the database", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := dish.NewInDBRepository(db)

		dishes, err := repo.GetAll()
		assert.NoError(err)

		assert.Equal(defaultDishes, dishes)
	})
}

func TestInDBDishRepo_Create(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should insert the given dish into the database", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := dish.NewInDBRepository(db)

		d := model.NewDish(4, "Dish 4")
		err = repo.Create(d)
		assert.NoError(err)

		dishes, err := repo.GetAll()
		assert.NoError(err)

		assert.Contains(dishes, d)
	})
}

func TestInDBDishRepo_Remove(t *testing.T) {
	test.SkipIntegration(t)

	assert := assertpkg.New(t)

	t.Run("should remote the given dish from the database", func(t *testing.T) {
		BeforeEach(t)
		defer AfterEach(t)

		db, err := openDB()
		assert.NoError(err)

		defer func(db *sql.DB) { _ = closeDB(db) }(db)

		repo := dish.NewInDBRepository(db)

		err = repo.Remove(3)
		assert.NoError(err)

		dishes, err := repo.GetAll()
		assert.NoError(err)

		expectedDishes := defaultDishes[0 : len(defaultDishes)-1]
		assert.Equal(expectedDishes, dishes)
	})
}
