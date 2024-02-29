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

func NewIngredient(ID int, name string) *Ingredient {
	return &Ingredient{
		ID:   ID,
		Name: name,
	}
}

func (i Ingredient) GetID() int {
	return i.ID
}
