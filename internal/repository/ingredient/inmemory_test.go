package ingredient_test

import (
	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/model"
	"platter/internal/repository"
	"platter/internal/repository/ingredient"
	"platter/internal/test"
	"testing"
)

func TestInMemIngredientRepo_New(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("should return an in-memory ingredient repository with no ingredients", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		ingredients, err := repo.GetAll()
		assert.Nil(err)
		assert.Empty(ingredients)
	})
}

func TestInMemIngredientRepo_Create(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("should add the given ingredient to the repository", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		i := model.NewIngredient(0, "test ingredient")

		err := repo.Create(i)
		assert.Nil(err)

		ingredients, err := repo.GetAll()
		assert.Nil(err)
		assert.NotEmpty(ingredients)
		assert.Contains(ingredients, i)
	})

	t.Run("shouldn't add the same ingredient twice", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		i := model.NewIngredient(0, "test ingredient")

		err := repo.Create(i)
		assert.Nil(err)

		err = repo.Create(i)
		assert.ErrorIs(err, repository.ErrAlreadyExists)
	})
}

func TestInMemIngredientRepo_Remove(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("shouldn't remove an ingredient if it does not exist", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()

		err := repo.Remove(1)
		assert.ErrorIs(err, repository.ErrNotExists)
	})

	t.Run("shouldn't remove an ingredient if it does not exist", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		i := model.NewIngredient(0, "test ingredient")

		err := repo.Create(i)
		assert.Nil(err)

		ingredients, err := repo.GetAll()
		assert.Nil(err)
		assert.Contains(ingredients, i)

		err = repo.Remove(i.GetID())
		assert.Nil(err)

		ingredients, err = repo.GetAll()
		assert.Nil(err)
		assert.NotContains(ingredients, i)
	})
}
