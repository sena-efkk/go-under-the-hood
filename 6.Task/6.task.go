package main

import "fmt"

func main() {
	//ygulama 1: Sayı kontrolü
	//
	//Aşağıdaki programı sen yaz:
	//
	//number isminde bir sayı oluştur.
	//Sayı pozitifse "Pozitif" yazdır.
	//Negatifse "Negatif" yazdır.
	//Sıfırsa "Sıfır" yazdır.
	//Ayrıca sayının çift veya tek olduğunu kontrol et.

	number := []int{5, 22, 6, 13, 2, 0, -15, -22}

	for _, num := range number {
		if num < 0 {
			fmt.Println("sayi negatiftir", num)
			if num%2 == -1 {
				fmt.Println("bu sayı tektir", num)
			} else if num%2 == 0 {
				fmt.Println("bu sayı çifttir", num)

			} else if num == 0 {
				fmt.Println("bu sayı sıfırdır.", num)
			}
		} else if num > 0 {
			fmt.Println("sayi pozitiftir", num)
			if num%2 == 1 {
				fmt.Println("bu sayı tektir", num)
			} else if num%2 == 0 {
				fmt.Println("bu sayı çifttir", num)

			}
		} else if num == 0 {
			fmt.Println("bu sayı sıfırdır.", num)
		}

	}
}
