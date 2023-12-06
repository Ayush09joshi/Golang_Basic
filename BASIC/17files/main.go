package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//We can only work with the text files in Golang (Wihout using some specific packages.)
	fmt.Println("Welcome to working with files in golang")

	content := "This is the content we wanna push."

	file, err := os.Create("./contentFile.txt") //Creating a file named:contentFile

	// if err != nil {
	// 	panic(err) --> Panic will end the exe of the program, and ere we area sking it to throw the error it face.
	// }

	checkNilErr(err)

	//Now we are going to write in the file we created.
	// io.WriteString(file, content) --> This will give you the length

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is: ", length)

	//Closing the file
	defer file.Close()

	readFile("./contentFile.txt")
}

//Now for reading te file.

func readFile(filename string) {

	//Data is always read in the byte and not in string or something else.
	dataByte, err := os.ReadFile(filename) //To read the file

	checkNilErr(err)

	//Wraping the value into the string, else we will get data in byte format like [30 70 ....]
	fmt.Println("The data in the File is:\n", string(dataByte))

}

// Creating a func for checking errors.
func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}

}
