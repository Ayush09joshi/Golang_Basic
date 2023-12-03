package main

import (
	"fmt"
	"sort"
)

func main() {

	//These are dynamic arrays in general.
	fmt.Println("Welcome to the Slices!!")

	//Way to declare slice.
	var fruits = []string{}
	fruits = append(fruits, "Mango", "Banana", "Kiwi", "Apple", "Peach")

	fmt.Println("The items in the slice are: ", fruits)
	fmt.Println("The items in the slice are: ", len(fruits))

	//Slicing the Slice.

	fruits = append(fruits[1:4]) //Here index 1 is inclusive and 3 is non inclusive.
	fmt.Println("The items in the slice, after slicing it are: ", fruits)
	fmt.Println("The items in the slice are: ", len(fruits))

	highscore := make([]int, 4)

	highscore[0] = 69
	highscore[1] = 6969
	highscore[2] = 696969
	highscore[3] = 69696969
	//highscore[4] = 69696969 --> This will give an out of bound error.
	//Since, the memory allocated toslice is full, we cannot add more data to it.

	//Addig more data in highscore.
	highscore = append(highscore, 420, 555, 619)

	fmt.Println("Data present in highscore is: ", highscore)

	//Sorting the values of highscore.
	sort.Ints(highscore)
	fmt.Println(highscore)

	//Checking if data in highscore is sort or not.
	fmt.Println("Is the data sorted? ", sort.IntsAreSorted(highscore))

	//Remove an element from slice based on the index.
	var courses = []string{"nextJS", "reactJS", "JS", "swift", "pyton"}
	fmt.Println(courses)

	//Removing Element
	var index int = 2 //Initializing which index have to be remove.

	//Slicing the Slice into 2 parts and Merge them Together.
	courses = append(courses[:index], courses[index+1:]...) //Removing element (***...*** IS IMP)
	fmt.Println(courses)

}
