package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer <= 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendList []string, myList []string) {
	secretIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	factor := float64(portions) / 2.0

	var scaledQuantities = make([]float64, len(quantities))
	_ = copy(scaledQuantities, quantities)

	for i := range scaledQuantities {
		scaledQuantities[i] *= factor
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
