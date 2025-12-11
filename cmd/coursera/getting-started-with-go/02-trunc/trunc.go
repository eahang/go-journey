package main

import (
	"fmt"
)

/*
*

	Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.
	Truncation is the process of removing the digits to the right of the decimal place.
*/
func main() {
	var num float64

	fmt.Print("Please enter a floating point number: ")

	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Error reading input. Please ensure you enter a valid number.")
		return
	}

	var nInt int64 = int64(num)

	fmt.Printf("The truncated integer is: %d\n", nInt)
}
