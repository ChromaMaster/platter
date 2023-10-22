package cooking_test

import (
	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/cooking"
	"testing"
)

func TestIngredient_New(t *testing.T) {
	assert := assertpkg.New(t)
	t.Run("should return an ingredient with the given name", func(t *testing.T) {
		ingredient := cooking.NewIngredient("test ingredient", cooking.IngredientTypeDairy)
		assert.Equal("test ingredient", ingredient.Name)
	})

	t.Run("should return an ingredient with the given type", func(t *testing.T) {
		ingredient := cooking.NewIngredient("test ingredient", cooking.IngredientTypeDairy)
		assert.Equal(cooking.IngredientTypeDairy, ingredient.Type)
	})
}

func TestIngredient_Equal(t *testing.T) {
	assert := assertpkg.New(t)
	t.Run("should return true if the given ingredient has the same name", func(t *testing.T) {
		ingredient := cooking.NewIngredient("test ingredient", cooking.IngredientTypeFruit)
		otherIngredient := cooking.NewIngredient("test ingredient", cooking.IngredientTypeLegume)

		assert.True(ingredient.Equal(otherIngredient))
	})

	t.Run("should return false if the given ingredient has a different name", func(t *testing.T) {
		ingredient := cooking.NewIngredient("test ingredient", cooking.IngredientTypeGrain)
		otherIngredient := cooking.NewIngredient("other test ingredient", cooking.IngredientTypeVegetable)

		assert.False(ingredient.Equal(otherIngredient))
	})
}
