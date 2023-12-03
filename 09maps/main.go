package main

import "fmt"

func main() {

	println("Welome to Maps in Goloang!")

	//Creating Map.
	languages := make(map[string]string)

	//Adding elements to Map.
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("Lists of all the languages are: ", languages)


	//Printing the value for individual key.
	fmt.Println("JS is the key for: ", languages["JS"])


	//Deleting the value in map.
	delete(languages, "RB") //This method can be used in Slices as well.
	fmt.Println("Lists of all the languages are: ", languages)


	//Basic idea of Loops in Maps.
	for key, value := range languages {
		fmt.Printf("For key %v value is %v.\n" , key, value)
	}

}