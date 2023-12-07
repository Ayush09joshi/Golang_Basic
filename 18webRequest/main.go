package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Welcome to handing web request.")

	//Get request.
	response, err := http.Get(url)

	//Checking for errors, if available.
	checkNilErr(err)

	fmt.Printf("Response is of type: %T", response)

	//Closing the connection request. (Alwasys Close Manually.)
	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println("Data in dataByte is: \n", string(databyte))

}

func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}
}
