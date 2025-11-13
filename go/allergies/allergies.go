package allergies

var allergenByName = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var allergyNames []string
	for allergenName, allergen := range allergenByName {
		if allergies&allergen == allergen {
			allergyNames = append(allergyNames, allergenName)
		}
	}
	return allergyNames
}

func AllergicTo(allergies uint, allergenName string) bool {
	allergen := allergenByName[allergenName]
	return allergen > 0 && allergies&allergen == allergen
}
