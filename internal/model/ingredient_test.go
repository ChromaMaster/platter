package model_test

import (
	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/model"
	"platter/internal/test"
	"testing"
)

func TestIngredient_New(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)
	t.Run("should return an ingredient with the given name", func(t *testing.T) {
		ingredient := model.NewIngredient("test ingredient")
		assert.Equal("test ingredient", ingredient.Name)
	})
}

func TestIngredient_GetID(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)
	t.Run("should return the ingredient's ID", func(t *testing.T) {
		ingredient := model.NewIngredient("test ingredient")
		assert.Equal(0, ingredient.GetID())
	})
}
