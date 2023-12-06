package main

import "fmt"

func main() {

	fmt.Println("Welcome to Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	//Way-1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//Way-2
	for i := range days { // range works as a loop itself and moves to all the element in slice and map.
		fmt.Println(days[i])
	}

	//For-each type of loop
	for index, days := range days { //index, days = INDEX, VALUE
		fmt.Printf("Index is %v and value is %v \n", index, days)
	}

	//Break, Continue and goto in golang
	newVal := 1

	for newVal < 10 {

		if newVal == 2 {
			goto JSR
		}

		if newVal == 5 {
			newVal++
			continue
			//break
		}

		fmt.Println("Value is: ", newVal)
		newVal++
	}

JSR:
	fmt.Println("Jai Shree Ram")
}
