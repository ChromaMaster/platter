package ingredient

import (
	"database/sql"
	"fmt"
	"platter/internal/model"
	"platter/internal/repository"
)

type Repository interface {
	repository.Repository[model.Ingredient]
}

type InDBIngredientRepository struct {
	db *sql.DB
}

func NewInDBIngredientRepository(db *sql.DB) *InDBIngredientRepository {
	return &InDBIngredientRepository{
		db: db,
	}
}

func (i InDBIngredientRepository) Create(ingredient *model.Ingredient) error {
	q := `INSERT INTO ingredients (name) VALUES (?)`
	stmt, err := i.db.Prepare(q)
	if err != nil {
		return fmt.Errorf("could not prepare the insert query: %w", err)
	}

	defer func(stmt *sql.Stmt) { _ = stmt.Close() }(stmt)

	r, err := stmt.Exec(ingredient.Name)
	if err != nil {
		return fmt.Errorf("could not execute the insert query: %w", err)
	}

	if num, err := r.RowsAffected(); err != nil || num != 1 {
		return fmt.Errorf("non-effective insert: %w", err)
	}

	return nil
}

func (i InDBIngredientRepository) GetAll() ([]*model.Ingredient, error) {
	q := `SELECT id, name FROM ingredients`
	rows, err := i.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("could not execute the query: %w", err)
	}

	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	ingredients := make([]*model.Ingredient, 0)

	for rows.Next() {
		i := &model.Ingredient{}

		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, fmt.Errorf("could not get the ingredient value: %w", err)
		}

		ingredients = append(ingredients, i)
	}

	return ingredients, nil
}

func (i InDBIngredientRepository) Remove(ID int) error {
	q := `DELETE FROM ingredients WHERE id = ?`
	stmt, err := i.db.Prepare(q)
	if err != nil {
		return fmt.Errorf("could not prepare the delete query %w", err)
	}

	defer func(stmt *sql.Stmt) { _ = stmt.Close() }(stmt)

	r, err := stmt.Exec(ID)
	if err != nil {
		return fmt.Errorf("could not exevure the delete query %w", err)
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return fmt.Errorf("delete query failed %w", err)
	}

	return nil
}
