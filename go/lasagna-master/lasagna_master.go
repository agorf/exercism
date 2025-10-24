package lasagna

func PreparationTime(layers []string, averageLayerPrepMinutes int) int {
	if averageLayerPrepMinutes == 0 {
		averageLayerPrepMinutes = 2 // Default
	}
	return len(layers) * averageLayerPrepMinutes
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		scaledQuantities[i] = (float64(numPortions) / 2.0) * quantity
	}
	return scaledQuantities
}
