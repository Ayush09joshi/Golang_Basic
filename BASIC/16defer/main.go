package main

import (
	"fmt"

)

func main() {

	//Defer push the line in which it is used at the last (for execution).
	defer fmt.Println("This is defer_1")
	defer fmt.Println("This is defer_2")
	defer fmt.Println("This is defer_3")
	//In case of multiple defer, it uses LIFO

	fmt.Println("Welcome to Deffer in golang")

	myDefer()


}

func myDefer(){

	for i := 0; i <5; i++ {
		defer fmt.Print("\n", i)
	}
}