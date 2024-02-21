package lasagna

import "fmt"

//defining what a lasagna should be [aka Interface]
type Lasagna struct {
	name   string
	layers int
	time   int
}

// TODO: define the 'OvenTime' constant
const OvenTime int = 40

func main() {
	//setting the Object
	lasagnaProperties := Lasagna{
		name:   "Big Boy Dish",
		layers: 4,
		time:   40,
	}

	//printing
	fmt.Print("The name of it is ", lasagnaProperties.name, "\nWith a prep time of ", PreparationTime(lasagnaProperties.layers), " minutes\n")
}

// Taking total prep time and removing the minutes that it's been there already
func RemainingOvenTime(actualMinutesInOven int) int {
	return (PreparationTime(20) - actualMinutesInOven)
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return (numberOfLayers * 2)
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (PreparationTime(numberOfLayers) + actualMinutesInOven)
}
