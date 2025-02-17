package lasagna

func PreparationTime(layers []string, t int) int {
	if t == 0 {
		t = 2
	}
	return len(layers) * t
}

func Quantities(layers []string) (noodles int, sauces float64) {
	noodles = 0
	sauces = 0
	for _, layer := range layers {
		if layer == "sauce" {
			sauces += 0.2
		}

		if layer == "noodles" {
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(friend []string, mine []string) {
	mine[len(mine)-1] = friend[len(friend)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	n := float64(portions) / 2
	newQuantities := make([]float64, len(quantities))
	for i, q := range quantities {
		newQuantities[i] = q * n
	}

	return newQuantities
}
