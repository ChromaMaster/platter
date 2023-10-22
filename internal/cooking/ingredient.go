package cooking

type IngredientType int

const (
	IngredientTypeDairy IngredientType = iota
	IngredientTypeMeatAndEggs
	IngredientTypeLegume
	IngredientTypeVegetable
	IngredientTypeFruit
	IngredientTypeGrain
	IngredientTypeFat
)

type Ingredient struct {
	Name string
	Type IngredientType
}

func NewIngredient(name string, ingredientType IngredientType) *Ingredient {
	return &Ingredient{
		Name: name,
		Type: ingredientType,
	}
}

func (i *Ingredient) Equal(otherIngredient *Ingredient) bool {
	return i.Name == otherIngredient.Name
}
