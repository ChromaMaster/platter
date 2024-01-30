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

func NewInDBIngredientRepository(db *sql.DB) *InDBIngredientRepository {
	return &InDBIngredientRepository{
		db: db,
	}
}

func (i InDBIngredientRepository) Create(model *model.Ingredient) error {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

type InDBIngredientRepository struct {
	db *sql.DB
}
