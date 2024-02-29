package model_test

import (
	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/model"
	"platter/internal/test"
	"testing"
)

func TestDish_New(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)
	t.Run("should return an dish with the given name", func(t *testing.T) {
		dish := model.NewDish(0, "test dish")
		assert.Equal("test dish", dish.Name)
	})

	t.Run("should return an dish with no ingredients", func(t *testing.T) {
		dish := model.NewDish(0, "test dish")
		assert.Equal(0, len(dish.Ingredients))
	})
}

func TestDish_GetID(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)
	t.Run("should return the dish's ID", func(t *testing.T) {
		dish := model.NewDish(2, "test dish")
		assert.Equal(2, dish.GetID())
	})
}
