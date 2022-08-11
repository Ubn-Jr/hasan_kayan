// Binary to decimal converter code
package main

import "fmt"

//func divides decimal number to 2 and add the reminder in to digits array until decimal cannot divided
// array has 64 digits, so we can calculate 64 digits of binary
func decimalToBinary(decimalNumber int) [64]int {
	var digits [64]int
	var digitNo int = 63
	for decimalNumber != 0 {
		var integer int = decimalNumber % 2
		decimalNumber = decimalNumber / 2
		digits[digitNo] = integer
		//In binary conversion, since units are written from the end, we try to print from the end in the array.
		digitNo--

	}
	// this section limits the range we want to print, otherwise all of array digits will be preinted even if we dont give it a value 
	fmt.Print("Binary equivalent of your number : ")
	for printRange := 63; printRange != digitNo; printRange-- {
		fmt.Print(digits[printRange])
	}
	return digits
}

func main() {
	var numDeci int
	fmt.Print("Enter the number you want to convert in binary :")
	fmt.Scan(&numDeci)
	decimalToBinary(numDeci)

}
