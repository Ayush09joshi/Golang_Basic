package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {


	//WAY-1
	welcome := "Hey! Welcome to the platform"
	fmt.Println(welcome)

	//Taking input from the user.
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for our Pizza: ")


	//COMMA OK || COMMA ERROR
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is: %T", input)




	//WAY-2
	welcome2 := "Hey! Welcome to the platform 2"
	fmt.Println(welcome2)
	var x int
	fmt.Println("Enter the rating for our Pizza 2: ")
	fmt.Scan(&x)
	fmt.Println("The rating for the Pizza 2 is: ", x)




}