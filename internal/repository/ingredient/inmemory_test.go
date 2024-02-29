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

	t.Run("should return an in-memory ingredient repository", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		assert.NotNil(repo)
	})
}

func TestInMemIngredientRepo_Init(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("should initialize the in-memory ingredient repository", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		err := repo.Init()
		assert.Nil(err)
	})

}

func TestInMemIngredientRepo_Create(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("should add the given ingredient to the repository", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		i := model.NewIngredient(1, "test ingredient")

		err := repo.Create(i)
		assert.Nil(err)

		ingredients, err := repo.GetAll()
		assert.Nil(err)
		assert.NotEmpty(ingredients)
		assert.Contains(ingredients, i)
	})

	t.Run("shouldn't add the same ingredient twice", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		i := model.NewIngredient(1, "test ingredient")

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

	t.Run("should remove an ingredient if it does exist", func(t *testing.T) {
		repo := ingredient.NewInMemIngredientRepository()
		i := model.NewIngredient(1, "test ingredient")

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
