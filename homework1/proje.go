package main

import "fmt"

func index(length float32, weigth float32) {
	dob := length * length
	ind := weigth / dob
	fmt.Print(ind)
}

func main() {
	var weigth float32
	var length float32
	fmt.Print("Boyunuzu giriniz metre cinsinden:\n")
	fmt.Scan(&length)
	fmt.Print("Kilonuzu giriniz kilogram cinsinden:\n")
	fmt.Scan(&weigth)
	fmt.Println("Vucut kitle indeksiniz:")
	index(length, weigth)

}
