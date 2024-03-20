//basic value printing

package values

import "fmt"

func main()  {
	fmt.Println("go"+"Lang")

	//simple strings
	var welcomer string = "Hey hey people"
	followUp := "Astro here ;)"

	fmt.Printf("%s, %v", welcomer, followUp)

	//arrays & slices
	names := [3]string{"Kallyadranoch", "Allihanna", "Lenah"}
	fmt.Println("Godhood:", names)

	gods := []string{"Tenebra", "Azgher", "Lin Wu", "Arsenal"}
	gods = append(gods, "Wynna")  //slice = append(slice, variable)
	fmt.Println("Divine:", gods)



	
}