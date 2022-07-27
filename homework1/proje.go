package main

import "fmt"

func index(leng float32, weig float32) {
	var ind float32
	var dob float32
	dob = leng * leng
	ind = weig / dob
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
