package lasagna

// TODO: define the 'PreparationTime()' function

/*
PreparationTime calculates the average preparation time needed
to prepare all lasagna's layers. The function receives
a []string variable with the lasagna's layers, and
an int variable, which value is the average preparation time per layer.
The function returns an int value, which is the total time needed
*/
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}


// TODO: define the 'Quantities()' function

/*
Quantities calculates the total noodles and sauce needed to make
the lasagna. The function receives a []string variable with the types of
each lasagna's layers. The function returns two values, (int, float64),
the int value is the grams of noodles needed and the float64 value is
the liters of sauce needed for the lasagna
*/
func Quantities(layers []string) (int, float64) {
	var gramsPerNoodlesLayer int = 50
	var litersPerSauceLayer float64 = 0.2
	var totalNoodlesNeeded int
	var totalSauceNeeded float64

	for i := 0; i < len(layers); i++ {
		currentLayer := layers[i]

		switch currentLayer {
		case "sauce":
			totalSauceNeeded += litersPerSauceLayer
		case "noodles":
			totalNoodlesNeeded += gramsPerNoodlesLayer
		}
	}

	return totalNoodlesNeeded, totalSauceNeeded
}

// TODO: define the 'AddSecretIngredient()' function

/*
AddSecretIngredient replaces each lasagna's layer where is expected to
put the secret ingredient, which is the final ingredient in the reference list.
The function receives two []string variables, the first parameter is
the reference list of ingredients, and the second parameter is the
list of layer to add the secret ingredient. The function replaces the layers
with the "?" string are replaced by the secret ingredient.
*/
func AddSecretIngredient(refList, editableList []string)  {
	var secretIngredient string = refList[ len(refList) - 1 ]

	for i := 0; i < len(editableList); i++ {
		if editableList[i] == "?" {
			editableList[i] = secretIngredient
		}
	}
}

// TODO: define the 'ScaleRecipe()' function

/*
ScaleRecipe calculates the amounts materials for the desired number of
portions. The function receives a []float64 as first parameter, which is
the quantity list of the recipes based on two portions, and a int parameter
which is the desired number of portions. The function returns a []float64
with the needed quantities for the desired number of portions
*/
func ScaleRecipe(originalQuantities []float64, desiredPortions int) []float64 {
	var scaleFactor float64 = float64(desiredPortions) / 2
	var newQuantities []float64 = make([]float64, len(originalQuantities))

	for i := 0; i < len(originalQuantities); i++ {
		newQuantities[i] = originalQuantities[i] * scaleFactor
	}

	return newQuantities
}
