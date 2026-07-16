package main

import "fmt"

func printAnimals(animals []string) {
	for index, animal := range animals {
		fmt.Printf("%d indexsi %s hyayvanım\n", index, animal)

	}
}
func isCalf(animal string) bool {
	return animal == "Buzağı"
}
func countAnimals(animals []string) int {
	return len(animals)
}

func countCalves(animals []string) int {
	//count := 0
	//for i := 0; i < len(animals); i++ {
	//	if animals[i] == "Buzağı" {
	//		count++
	//	}
	//}
	//return count
	count := 0

	for _, animal := range animals {
		if isCalf(animal) {
			count++
		}
	}

	return count
}

func findAnimal(animals []string, target string) (string, bool) {
	for _, animal := range animals {
		if animal == target {
			return animal, true
		}
	}
	return "", false
}

// belirli hayvan turunu sayma fonksiyonu
func countByType(animals []string, target string) int {
	total := 0
	for _, animal := range animals {
		if animal == target {
			total++
		}
	}
	return total
}

// Hayvanın var olup olmadığını kontrol etme
func containsAnimal(animals []string, target string) bool {
	for _, animal := range animals {
		if animal == target {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{
		"kedi",
		"köpek",
		"ördek",
		"Buzağı",
	}

	ani, found := findAnimal(animals, "maymun")
	if found {
		fmt.Println("bulunan hayvan :", ani)
	} else {
		fmt.Println(" hayvanı bulunamadı.")
	}

	if containsAnimal(animals, "köpek") {
		fmt.Println("Listede köpek bulunuyor.")
	}

	printAnimals(animals)

	result := isCalf("Buzağı")
	fmt.Println(result)

	if isCalf("Buzağı") {
		fmt.Println("yaş kontrolu yapılmalıdır.")
	}

	fmt.Println("toplam hayvan sayısı: ", countAnimals(animals))
	fmt.Println("buzagı sayısı ektedir: ", countCalves(animals))
}
