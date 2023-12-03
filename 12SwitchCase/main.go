package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Welcome to Switch-Case in golang")

	//Explanation through ludo game.

	//Creating a Random number for die.

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Die is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Die value is 1, move 1 spot ahead.")
	case 2:
		fmt.Println("Die value is 2, move 2 spot ahead.")
	case 3:
		fmt.Println("Die value is 3, move 3 spot ahead.")
	case 4:
		fmt.Println("Die value is 4, move 4 spot ahead.")
	case 5:
		fmt.Println("Die value is 5, move 5 spot ahead.")
	case 6:
		fmt.Println("Die value is 6, Open, move 6 spot ahead.")
		fallthrough //When fallthrough is called, next argument will run with it.
	default:
		fmt.Println("Roll the Dice again.")

	}

}
