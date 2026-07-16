package main

import "fmt"

func main() {
	animals2 := []string{
		"inek",
		"kedi",
		"aslan",
	}

	for index, animal := range animals2 {
		fmt.Printf("%d hayvanımın indexi %s", index+1, animal)

		if animal == "inek" {
			println("bu hayv		an inektir")
		} else if animal == "kedi" {
			println("bu hayvan kedidir")
		} else {
			println("bu hayvan aslandır")
		}

		switch animal {
		case "aslan", "kedi":
			fmt.Println("kedigiller ailesi")
		case "inek":
			fmt.Println("bu inektie")
		default:
			fmt.Println("Bilinmeyen hayvan türü")
		}

	}

	//Liste üzerinde dön.
	//
	//Her hayvan için:
	//
	//1 - İnek
	//2 - Boğa
	//3 - Buzağı
	//4 - Düve
	//
	//şeklinde çıktı üret.
	//
	//Buzağı geldiğinde ayrıca:
	//
	//Bu hayvan için yaş kontrolü yapılmalıdır.
	//
	//yazdır.
	animals := []string{
		"İnek",
		"Boğa",
		"Buzağı",
		"Düve",
	}
	for index, animal := range animals {
		fmt.Printf("%d - %s\n", index+1, animal)

		if animal == "Buzağı" {
			fmt.Println("Bu hayvan için yaş kontrolü yapılmalıdır.")
		}
	}

}
