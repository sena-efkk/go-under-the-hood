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

	yas := 20
	rol := "admin"

	if yas > 18 && rol == "admin" {
		fmt.Println("Tam erişim sağlandı")
	} else if yas > 18 && rol == "user" {
		fmt.Println("Sınırlı erişim")
	} else {
		fmt.Println("Erişim reddedildi.")
	}
	if puan := 85; puan > 80 {
		fmt.Println("Bonus puan kazandınız")
		//puan değişkeni o süslü parantezlerin ({ })
		//içinde doğar ve kapanış parantezinde ölür.
		//Sen dışarıdan ona seslendiğinde, derleyici
		//"Böyle biri hiç yaşamadı" der
	}

}
