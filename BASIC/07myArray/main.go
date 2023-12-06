package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	//Way-1 of declaring the Array.
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Kiwi"

	fmt.Println("Elements in the array are: ", fruits)
	fmt.Println("Elements in the array are: ", len(fruits))

	//Way-2 of declaring the Array.

	var vegie = [3]string{"Tomato", "Potato", "Garlic"}

	fmt.Println("Elements in the array are: ", vegie)
	fmt.Println("Elements in the array are: ", len(vegie))

}
