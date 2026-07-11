package main

import "fmt"

func main() {
	//10 tane doğal sayı yazan programı yazınız
	for i := 1; i < 11; i++ {
		fmt.Print(i)
	}

	//ilk 10 doğal sayının toplamını ekrana yazdıran program
	toplam := 0
	for j := 1; j < 11; j++ {
		toplam = toplam + j

	}
	fmt.Println()
	fmt.Print(toplam)

}
