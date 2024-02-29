package dish

import (
	"database/sql"
	"fmt"
	"platter/internal/model"
	"platter/internal/repository"
)

type Repository interface {
	repository.Repository[model.Dish]
}

type InDBRepository struct {
	db *sql.DB
}

func NewInDBRepository(db *sql.DB) *InDBRepository {
	return &InDBRepository{
		db: db,
	}
}

func (i InDBRepository) Init() error {
	q := `CREATE TABLE IF NOT EXISTS dishes (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		name VARCHAR(64) NULL
		);`

	if _, err := i.db.Exec(q); err != nil {
		return fmt.Errorf("could not create the table: %w", err)
	}

	return nil
}

func (i InDBRepository) Create(dish *model.Dish) error {
	q := `INSERT INTO dishes (name) VALUES (?)`
	stmt, err := i.db.Prepare(q)
	if err != nil {
		return fmt.Errorf("could not prepare the insert query: %w", err)
	}

	defer func(stmt *sql.Stmt) { _ = stmt.Close() }(stmt)

	r, err := stmt.Exec(dish.Name)
	if err != nil {
		return fmt.Errorf("could not execute the insert query: %w", err)
	}

	if num, err := r.RowsAffected(); err != nil || num != 1 {
		return fmt.Errorf("non-effective insert: %w", err)
	}

	return nil
}

func (i InDBRepository) GetAll() ([]*model.Dish, error) {
	q := `SELECT id, name FROM dishes`
	rows, err := i.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("could not execute the query: %w", err)
	}

	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	dishes := make([]*model.Dish, 0)

	for rows.Next() {
		d := model.NewDish(0, "")

		if err := rows.Scan(&d.ID, &d.Name); err != nil {
			return nil, fmt.Errorf("could not get the dish value: %w", err)
		}

		dishes = append(dishes, d)
	}

	return dishes, nil
}

func (i InDBRepository) Remove(ID int) error {
	q := `DELETE FROM dishes WHERE id = ?`
	stmt, err := i.db.Prepare(q)
	if err != nil {
		return fmt.Errorf("could not prepare the delete query %w", err)
	}

	defer func(stmt *sql.Stmt) { _ = stmt.Close() }(stmt)

	r, err := stmt.Exec(ID)
	if err != nil {
		return fmt.Errorf("could not execure the delete query %w", err)
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return fmt.Errorf("delete query failed %w", err)
	}

	return nil
}
