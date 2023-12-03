package main

import "fmt"

const name string = "Ayush09joshi" // --> Private Variable "n" in name is **small n**
const Name string = "Ayush02joshi" // ---> Public Variable "n" in name is **large n**

func main() {

	//Assigning values and printing their type.
	var name string = "Ayush"
	fmt.Println(name)
	fmt.Printf("Type of variable is: %T \n", name)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable is: %T \n", isLoggedIn)



	//Default values and some aliases.
	var sampleString string
	fmt.Println(sampleString)
	fmt.Printf("Type of variable is: %T \n", sampleString)


	//Implicit Type
	var website = "www.google.co.in" //NO TYPE IS DEFINED!!
	fmt.Println(website)
	fmt.Printf("Type of variable is: %T \n", website)


	//No var style
	novar := "no var declared"
	fmt.Println(novar) //Good Practice
	fmt.Printf("Type of variable is: %T \n", novar)
	println(novar) //Bad Practice

	//Calling private and public variable
	// fmt.Println(name) --> Gives Error as name is Private.
	fmt.Println(Name)	// --> Does not gives any error as it is Public.
}