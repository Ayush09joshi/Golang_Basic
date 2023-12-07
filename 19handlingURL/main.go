package main

import (
	"fmt"
	"net/url"
)

// Declaring a url we will work on.
// const myurl string = "https://lco.dev"
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {

	fmt.Println("Welcome to Handing the URLs in golang")
	fmt.Println(myurl)

	//Parsing the URL
	result, err := url.Parse(myurl)

	checkNilErr(err)

	fmt.Println("Schema: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())       //It is a Method.
	fmt.Println("RawQuery: ", result.RawQuery) //These are same as parameters in URL.

	//Working with quaryParameters.
	qparam := result.Query() //Quary() --> Gives a better way to read these parameters.

	//qparam have a value of url.Values [ie -> KEY VALUE PAIR]
	fmt.Printf("Type of qpram is: %T \n", qparam)

	//All key value pairs are of type String
	fmt.Println(qparam["coursename"])


	//Making URL from the chunks

	partsOfUrl := &url.URL{ //ALWAYS USE **&**
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
	}

	consURL := partsOfUrl.String()

	fmt.Println(consURL)

}

func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}
}
