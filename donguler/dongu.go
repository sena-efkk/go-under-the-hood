package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Println(" ")

	a := 0
	for a < 5 {
		fmt.Print(a)
		a++
	}
	fmt.Println()

	sayilar := []int{10, 20, 30}

	for index, deger := range sayilar {
		fmt.Printf("İndis: %d, Değer: %d\n", index, deger)
	}
}
