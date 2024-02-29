package model

type Dish struct {
	ID          int
	Name        string
	Ingredients []*Ingredient
}

func NewDish(ID int, name string) *Dish {
	return &Dish{
		ID:          ID,
		Name:        name,
		Ingredients: make([]*Ingredient, 0),
	}
}

func (d Dish) GetID() int {
	return d.ID
}
