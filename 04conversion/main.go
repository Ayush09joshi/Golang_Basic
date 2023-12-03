package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

	fmt.Println("Welcome to our Pizza App!")
	fmt.Printf("Enter the rating for our Pizza between 1 to 5: ")


	//Taking Input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	
	fmt.Println("Thanks for Rating: ", input)


	//Conversion of string to number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Added 1 to your rating: ", numRating + 1) //Adding 1 to the rating input.
	}
}