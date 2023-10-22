package cooking

type Ingredient struct {
	Name string
}

func (i *Ingredient) Equal(otherIngredient *Ingredient) bool {
	return i.Name == otherIngredient.Name
}
