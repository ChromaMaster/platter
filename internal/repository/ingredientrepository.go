package repository

import "platter/internal/model"

type IngredientRepository interface {
	Repository[model.Ingredient]
}

type InMemIngredientRepository struct {
	ingredients []*model.Ingredient
}

func NewInMemIngredientRepository() *InMemIngredientRepository {
	return &InMemIngredientRepository{}
}

func (i *InMemIngredientRepository) Create(ingredient *model.Ingredient) error {
	if i.contains(ingredient) {
		return ErrAlreadyExists
	}

	i.ingredients = append(i.ingredients, ingredient)

	return nil
}

func (i *InMemIngredientRepository) GetIngredients() ([]*model.Ingredient, error) {
	return i.ingredients, nil
}

func (i *InMemIngredientRepository) contains(ingredient *model.Ingredient) bool {
	for _, i := range i.ingredients {
		if i.GetID() == ingredient.GetID() {
			return true
		}
	}

	return false
}
