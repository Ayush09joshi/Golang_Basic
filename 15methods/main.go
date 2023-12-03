package main

import (
	"fmt"
)

func main() {

	//Most of the code is SAME code as Struct.

	fmt.Println("Welcome to Methods in golang")

	//Adding data to our Struct
	User1 := User{"Ayush", "joshi09ayush@gmail.com", true, 21}
	fmt.Println(User1)

	//For more clear data printing.
	fmt.Printf("Data of User1 is: %+v \n", User1)

	//For printing any specific data.
	fmt.Printf("Name and age of User1 is %v and %v respectively. \n", User1.Name, User1.Age)

	//Calling Method
	User1.GetStatus()

	User1.NewMail()

}

// Creating Structs
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Creating a Method:
func (u User) GetStatus() { //Capital G in GetStatus to keep it accessable.

	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {

	u.Email = "abc@xyz.com"
	fmt.Println("This is the new mail of the user: ", u.Email)

}
