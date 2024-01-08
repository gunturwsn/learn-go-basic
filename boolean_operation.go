package main

import "fmt"

func main() {
	var finalScore = 90
	var absence = 80

	var passFinalScore bool = finalScore > 80
	var passAbsence bool = absence > 80

	var pass bool = passFinalScore && passAbsence
	fmt.Println(pass)

}
