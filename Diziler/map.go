package main

import "fmt"

func main() {

	prices := map[string]float64{
		"USD": 123.456,
		"süt": 32.75,
	}

	for product, price := range prices {
		fmt.Printf("%s %f TL\n ", product, price)
	}

	text := "Go"

	for index, char := range text {
		fmt.Println(index, char)
	}
}
