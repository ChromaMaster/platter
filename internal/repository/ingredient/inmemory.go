package ingredient

import (
	"platter/internal/model"
	"platter/internal/repository"
)

type InMemRepository struct {
	ingredients []*model.Ingredient
}

func NewInMemIngredientRepository() *InMemRepository {
	return &InMemRepository{}
}

func (i *InMemRepository) Init() error {
	i.ingredients = make([]*model.Ingredient, 0)

	return nil
}

func (i *InMemRepository) Create(ingredient *model.Ingredient) error {
	if i.contains(ingredient) {
		return repository.ErrAlreadyExists
	}

	ingredient.ID = len(i.ingredients) + 1

	i.ingredients = append(i.ingredients, ingredient)

	return nil
}

func (i *InMemRepository) GetAll() ([]*model.Ingredient, error) {
	return i.ingredients, nil
}

func (i *InMemRepository) Remove(ID int) error {
	if !i.contains(&model.Ingredient{ID: ID}) {
		return repository.ErrNotExists
	}

	i.removeIngredientByID(ID)

	return nil
}

func (i *InMemRepository) contains(ingredient *model.Ingredient) bool {
	for _, i := range i.ingredients {
		if i.GetID() == ingredient.GetID() {
			return true
		}
	}

	return false
}

func (i *InMemRepository) removeIngredientByID(ID int) {
	for index, ing := range i.ingredients {
		if ing.GetID() == ID {
			i.ingredients = append(i.ingredients[:index], i.ingredients[index+1:]...)
		}
	}
}
