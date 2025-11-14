package allergies

var list = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(mask uint) []string {
	var result []string

	for k, v := range list {
		if mask&v != 0 {
			result = append(result, k)
		}
	}

	return result
}

func AllergicTo(mask uint, allergen string) bool {
	return mask&list[allergen] != 0
}
