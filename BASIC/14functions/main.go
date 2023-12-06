package main

import "fmt"

func main() {

	//Functions under a function is not allowed.

	fmt.Println("Welcome to funtions")

	//Function-1 Called
	greet()

	//Function-2 Called
	ans := adder(6, 9)
	fmt.Println("Result is: ", ans)

	//Function-3 Called
	proAns := proAdder(6, 9, 4, 2, 0)
	fmt.Println("Pro-Result is: ", proAns)

	//Function-4 Called

	//Way-1
	proAns2, _ := proAdder2(6, 9, 4, 2, 0)
	fmt.Println("Pro-Result is: ", proAns2)

	//Way-2
	proAns2, message := proAdder2(6, 9, 4, 2, 0)
	fmt.Println("Pro-Result is: ", proAns2, message)
}

//Simple Function (Function-1)
func greet() {
	fmt.Println("Welcome to the greet funtion!")
}

//Funtion with passed on values (Function-2)

//General Syntax --> func funName(name **type) returnType {}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

//Functin with dynamic input (Function-3)
func proAdder(values ...int) int { //Tells all values are of type int.

	total := 0
	for _, val := range values {
		total += val
	}

	return total
}

//Functin with dynamic input and multiple return values (Function-4)
func proAdder2(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Printing Function-4"
}
