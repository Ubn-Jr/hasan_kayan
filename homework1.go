package main

import "fmt"

// index function calculates user's body mass index according to  length and weigth
func index(leng float32, weig float32) {
	var ind, dob float32
	dob = leng * leng // length square
	ind = weig / dob
	fmt.Println(ind)

}

// fet_percent calculates user's fat rate as percent
func fat_percent(bel float32, age float32, gender float32) float32 {

	if gender == 1 {
		var percent float32 = 0.567*bel + 0.101*age - 31.8 // male users
		fmt.Println("Yağ yüzdeniz: %")
		fmt.Println(percent)
		return percent

	}
	if gender == 0 {
		var percent2 float32 = 0.439*bel + 0.221*age - 9.4 // kfemale users
		fmt.Println("Yağ yüzdeniz: %")
		fmt.Print(percent2)
		return percent2

	}
	return gender
}

func main() {
	var bely float32
	var age float32
	var gender float32

	var weigth float32
	var length float32
	// taking variables of user's to calculate mass index
	fmt.Print("Boyunuzu giriniz metre cinsinden:\n")
	fmt.Scan(&length)
	fmt.Print("Kilonuzu giriniz kilogram cinsinden:\n")
	fmt.Scan(&weigth)
	fmt.Println("Vucut kitle indeksiniz:")
	// calling index function by using user's variables
	index(length, weigth)
	// taking user's gender and other dimensions to calculate fat rate
	fmt.Printf("Cinsiyetiniz ERKEK ise '1' e basiniz KADIN ise '0' basiniz:\n")
	fmt.Scan(&gender)
	fmt.Printf("Bel cevrenizi giriniz(cm):\n")
	fmt.Scan(&bely)
	fmt.Printf("Yasinizi giriniz:\n")
	fmt.Scan(&age)
	// calling fat_percent function with user's variables
	fat_percent(bely, age, gender)
	// defining percent_fat fuınction output as a new variable
	var percent_fat float32 = fat_percent(bel, age, gender)
	// using percent_fat variable to give advices to user with conditions
	if percent_fat > 27 {
		fmt.Println("Yağ oranınız olması gerekenin üstünde.")
		fmt.Println("Karbonhidrat ve yağ alımınızı minimuma indirmeniz ve gereken beslenme listenizi mutlaka alın.")
	}
	if percent_fat < 11 && percent_fat > 8 {

		fmt.Println("Galiba sporla uğraşıyorsunuz, unutmayın belli bir yağ oranının altı sağlığüınız için zararlıdır, başarılar.")

		fmt.Println("Ortalama yağ oranının altındasınız, klilo almaya yönelik bir diyet düşünebilirsiniz.")
	}
	if percent_fat > 14 && percent_fat < 30 {

		fmt.Println("Olması gerekn kilo oranındasınız, formunuzu korumaya özen  gösterin.")

	}

}
