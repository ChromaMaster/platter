package cooking

import (
	"errors"
)

var ErrIngredientAlreadyAdded = errors.New("ingredient already added")

type Dish struct {
	Name        string
	Ingredients []*Ingredient
}

func NewDish(name string) *Dish {
	return &Dish{
		Name: name,
	}
}

func (d *Dish) AddIngredient(ingredient *Ingredient) error {
	if d.containsIngredient(ingredient) {
		return ErrIngredientAlreadyAdded
	}

	d.Ingredients = append(d.Ingredients, ingredient)

	return nil
}

func (d *Dish) containsIngredient(ingredient *Ingredient) bool {
	for _, i := range d.Ingredients {
		if i.Equal(ingredient) {
			return true
		}
	}

	return false
}
