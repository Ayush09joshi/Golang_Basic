package main

import (
	"fmt"
)

func main() {

	//BASICS OF STRUCTS
	//Structs in GO are the alternative to Classes in other languages.

	fmt.Println("Welcome to Structs in golang")

	//Adding data to our Struct
	User1 := User{"Ayush", "joshi09ayush@gmail.com", true, 21}
	fmt.Println(User1)

	//For more clear data printing.
	fmt.Printf("Data of User1 is: %+v \n", User1)

	//For printing any specific data.
	fmt.Printf("Name and age of User1 is %v and %v respectively.", User1.Name, User1.Age)

}

//Creating Structs

type User struct { //Type User -> It means that we are creating a new type "User" (int is a type, char is a type.)
	Name   string
	Email  string
	Status bool
	Age    int
}
