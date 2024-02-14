package main

import (
	"Top-University/universities"
	"fmt"
)

func main() {
	firstCity := "Austin"
	secondCity := "Perth"
	report := universities.NewReport()
	result, err := report.HighestInternationalStudents(firstCity, secondCity)
	if err != nil {
		panic("invalid response returned from the API..")
	}
	fmt.Println(result)
}
