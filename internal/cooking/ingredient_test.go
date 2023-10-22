package cooking_test

import (
	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/cooking"
	"testing"
)

func TestIngredient_Equal(t *testing.T) {
	assert := assertpkg.New(t)
	t.Run("should return true if the given ingredient has the same name", func(t *testing.T) {
		ingredient := &cooking.Ingredient{Name: "test ingredient"}
		otherIngredient := &cooking.Ingredient{Name: "test ingredient"}

		assert.True(ingredient.Equal(otherIngredient))
	})

	t.Run("should return false if the given ingredient has a different name", func(t *testing.T) {
		ingredient := &cooking.Ingredient{Name: "test ingredient"}
		otherIngredient := &cooking.Ingredient{Name: "other test ingredient"}

		assert.False(ingredient.Equal(otherIngredient))
	})
}
