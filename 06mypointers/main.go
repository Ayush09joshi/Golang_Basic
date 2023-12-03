package main

import "fmt"

func main() {

	fmt.Println("Welcome to the POINTERS!!")

	myNum := 23
	var ptr = &myNum

	fmt.Println("Value of prt is: ", ptr)
	fmt.Println("Value of prt is: ", *ptr)


}