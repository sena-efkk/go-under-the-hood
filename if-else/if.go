package main

import "fmt"

func main() {
	age := 2

	fmt.Println(age > 19)
	fmt.Println(age <= 12)

	if age < 12 {
		fmt.Println("yaşınız 12 dan kucuktur")

	} else if age == 12 {
		fmt.Println("yaşınız 12e eşittir")
	} else {
		fmt.Println("yaşınız 12 den buyktur")

	}

}
