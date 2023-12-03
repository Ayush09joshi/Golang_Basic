package main

import "fmt"

func main() {

	fmt.Println("Welcome to if_else ladder in golang")

	loginCout := 69
	var result string

	if loginCout < 10 {
		result = "Not a Regular User"
	} else if loginCout > 10 {
		result = "Regular User"
	} else {
		result = "Common User"
	}
	fmt.Println(result)

	//Initalising the value on the go in if_else ladder
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is more than 10")
	}

}
