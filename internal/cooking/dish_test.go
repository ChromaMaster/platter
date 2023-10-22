package cooking_test

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/cooking"
)

func TestDish_New(t *testing.T) {
	assert := assertpkg.New(t)
	t.Run("should return a dish with the given name", func(t *testing.T) {
		dish := cooking.NewDish("test dish")
		assert.Equal("test dish", dish.Name)
	})

	t.Run("should return a dish with no ingredients", func(t *testing.T) {
		dish := cooking.NewDish("test dish")
		assert.Empty(dish.Ingredients)
	})
}

func TestDish_AddIngredient(t *testing.T) {
	assert := assertpkg.New(t)
	t.Run("should add the given ingredient to the dish", func(t *testing.T) {
		dish := cooking.NewDish("test dish")
		ingredient := &cooking.Ingredient{Name: "test ingredient"}

		err := dish.AddIngredient(ingredient)
		assert.Nil(err)
		assert.NotEmpty(dish.Ingredients)
		assert.Contains(dish.Ingredients, ingredient)
	})

	t.Run("should not add the same ingredient twice", func(t *testing.T) {
		dish := cooking.NewDish("test dish")
		ingredient := &cooking.Ingredient{Name: "test ingredient"}

		err := dish.AddIngredient(ingredient)
		assert.Nil(err)

		err = dish.AddIngredient(ingredient)
		assert.ErrorIs(err, cooking.ErrIngredientAlreadyAdded)
	})
}
