package repository_test

import (
	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/model"
	"platter/internal/repository"
	"testing"
)

func TestInMemIngredientRepo_New(t *testing.T) {
	assert := assertpkg.New(t)
	t.Run("should return an in-memory ingredient repository with no ingredients", func(t *testing.T) {
		repo := repository.NewInMemIngredientRepository()
		ingredients, err := repo.GetIngredients()
		assert.Nil(err)
		assert.Empty(ingredients)
	})
}

func TestInMemIngredientRepo_Create(t *testing.T) {
	assert := assertpkg.New(t)
	t.Run("should add the given ingredient to the repository", func(t *testing.T) {
		repo := repository.NewInMemIngredientRepository()
		ingredient := model.NewIngredient(0, "test ingredient")

		err := repo.Create(ingredient)
		assert.Nil(err)

		ingredients, err := repo.GetIngredients()
		assert.Nil(err)
		assert.NotEmpty(ingredients)
		assert.Contains(ingredients, ingredient)
	})

	t.Run("shouldn't add the same ingredient twice", func(t *testing.T) {
		repo := repository.NewInMemIngredientRepository()
		ingredient := model.NewIngredient(0, "test ingredient")

		err := repo.Create(ingredient)
		assert.Nil(err)

		err = repo.Create(ingredient)
		assert.ErrorIs(err, repository.ErrAlreadyExists)
	})
}
