package dish

import (
	"platter/internal/model"
	"platter/internal/repository"
)

type InMemRepository struct {
	dishes []*model.Dish
}

func NewInMemRepository() *InMemRepository {
	return &InMemRepository{}
}

func (i *InMemRepository) Init() error {
	i.dishes = make([]*model.Dish, 0)

	return nil
}

func (i *InMemRepository) Create(dish *model.Dish) error {
	if i.contains(dish) {
		return repository.ErrAlreadyExists
	}

	dish.ID = len(i.dishes) + 1

	i.dishes = append(i.dishes, dish)

	return nil
}

func (i *InMemRepository) GetAll() ([]*model.Dish, error) {
	return i.dishes, nil
}

func (i *InMemRepository) Remove(ID int) error {
	if !i.contains(&model.Dish{ID: ID}) {
		return repository.ErrNotExists
	}

	i.removeDishByID(ID)

	return nil
}

func (i *InMemRepository) contains(dish *model.Dish) bool {
	for _, d := range i.dishes {
		if d.GetID() == dish.GetID() {
			return true
		}
	}

	return false
}

func (i *InMemRepository) removeDishByID(ID int) {
	for index, ing := range i.dishes {
		if ing.GetID() == ID {
			i.dishes = append(i.dishes[:index], i.dishes[index+1:]...)
		}
	}
}
