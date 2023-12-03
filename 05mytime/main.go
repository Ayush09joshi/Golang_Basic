package main

import (
	"fmt"
	"time"

)

func main() {
	fmt.Println("Welcome to time study of Golang")

	presentTime := time.Now()

	//Format for printing Date(DD/MM/YYY), Time(HH/MM/SS) and Day
	fmt.Println(presentTime.Format("02-01-2006 15-04-05 Monday"))


	//Creating Date Manually
	createDate := time.Date(2020, time.February, 6, 9, 9, 9, 9, time.UTC)
	fmt.Println(createDate)
}