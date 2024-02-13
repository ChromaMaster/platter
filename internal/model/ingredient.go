package model

type IngredientType int

const (
// IngredientTypeDairy IngredientType = iota
// IngredientTypeMeatAndEggs
// IngredientTypeLegume
// IngredientTypeVegetable
// IngredientTypeFruit
// IngredientTypeGrain
// IngredientTypeFat
)

type Ingredient struct {
	ID   int
	Name string
}

func NewIngredient(name string) *Ingredient {
	return &Ingredient{
		Name: name,
	}
}

func (i Ingredient) GetID() int {
	return i.ID
}
