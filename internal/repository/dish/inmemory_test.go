package dish_test

import (
	assertpkg "github.com/stretchr/testify/assert"
	"platter/internal/model"
	"platter/internal/repository"
	"platter/internal/repository/dish"
	"platter/internal/test"
	"testing"
)

func TestInMemDishRepo_New(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("should return an in-memory dish repository", func(t *testing.T) {
		repo := dish.NewInMemRepository()
		assert.NotNil(repo)
	})
}

func TestInMemIDishRepo_Init(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("should initialize the in-memory dish repository", func(t *testing.T) {
		repo := dish.NewInMemRepository()
		err := repo.Init()
		assert.Nil(err)
	})

}

func TestInMemDishRepo_Create(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("should add the given dish to the repository", func(t *testing.T) {
		repo := dish.NewInMemRepository()
		d := model.NewDish(1, "test dish")

		err := repo.Create(d)
		assert.Nil(err)

		dishes, err := repo.GetAll()
		assert.Nil(err)
		assert.NotEmpty(dishes)
		assert.Contains(dishes, d)
	})

	t.Run("shouldn't add the same dish twice", func(t *testing.T) {
		repo := dish.NewInMemRepository()
		d := model.NewDish(1, "test dish")

		err := repo.Create(d)
		assert.Nil(err)

		err = repo.Create(d)
		assert.ErrorIs(err, repository.ErrAlreadyExists)
	})
}

func TestInMemDishRepo_Remove(t *testing.T) {
	test.SkipUnit(t)

	assert := assertpkg.New(t)

	t.Run("shouldn't remove a dish if it does not exist", func(t *testing.T) {
		repo := dish.NewInMemRepository()

		err := repo.Remove(1)
		assert.ErrorIs(err, repository.ErrNotExists)
	})

	t.Run("should remove a dish if it does exist", func(t *testing.T) {
		repo := dish.NewInMemRepository()

		d := model.NewDish(1, "test dish")

		err := repo.Create(d)
		assert.Nil(err)

		dishes, err := repo.GetAll()
		assert.Nil(err)
		assert.Contains(dishes, d)

		err = repo.Remove(d.GetID())
		assert.Nil(err)

		dishes, err = repo.GetAll()
		assert.Nil(err)
		assert.NotContains(dishes, d)
	})
}
