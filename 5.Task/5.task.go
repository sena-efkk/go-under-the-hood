package main

import "fmt"

func main() {
	//Uygulama 2: Not sistemi
	//
	//Bir score değişkeni oluştur.
	//
	//90-100: A
	//80-89: B
	//70-79: C
	//60-69: D
	//0-59: F
	//0 altı veya 100 üstü: geçersiz not
	//
	//Bunu önce if / else, sonra switch ile yaz.
	var score int
	fmt.Print("bir değer giriniz:")
	fmt.Scanf("%d", &score)
	fmt.Println(score)

	//önce if else
	if score < 0 {
		fmt.Println("bu not sıfırdan az olamaz")
	} else if score > 100 {
		fmt.Println("bu not 100 den fazla olamaz")
	} else if 90 <= score && score <= 100 {
		fmt.Println("harf notu A")
	} else if 80 <= score && score <= 89 {
		fmt.Println("harf notu B")
	} else if 70 <= score && score <= 79 {
		fmt.Println("harf notu C")
	} else if 60 <= score && score <= 69 {
		fmt.Println("harf notu D")
	} else {
		fmt.Println("harf notu F")
	}

	//sonra switch case
}
