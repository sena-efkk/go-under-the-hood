package main

import "fmt"

func printAnimals(animals []string) {
	for index, animal := range animals {

		fmt.Printf("Animal %d : %s\n", index+1, animal)
		//Hayvanları numaralı yazdır.
		//Kullanıcıya gösterilen sıra 1 ile başlamalı.
	}
}

func countAnimals(animals []string) int {
	return len(animals)
	//Toplam hayvan sayısını döndür.
}

func isCalf(animal string) bool {
	return animal == "Buzağı"
	//Hayvan "Buzağı" ise true döndür

	//daha sade hali kısaca tek bunu yazabilirsin: return animal == "Buzağı"
}

func countCalves(animals []string) int {
	//isCalf() fonksiyonunu kullan.
	//Toplam buzağı sayısını döndür.
	count := 0
	for _, animal := range animals {
		if isCalf(animal) {
			count++
		}
	}
	return count
}

func containsAnimal(animals []string, target string) bool {
	//Aranan tür listede varsa true.
	//Hayvan bulunduğu anda erken return kullan. :erken return ne demek bilmiyorum ?
	for _, animal := range animals {
		if animal == target {
			return true
		}
	}
	return false
}

func countByType(animals []string, target string) int {
	//İstenen hayvan türünün kaç kez geçtiğini döndür.
	count := 0
	for _, animal := range animals {
		if animal == target {
			count++
		}
	}
	return count
}

func calculateTotalWeight(weights []float64) float64 {
	//Bütün ağırlıkları topla.
	total := 0.0
	for _, weight := range weights {
		total += weight
	}
	return total
}

func calculateAverageWeight(weights []float64) float64 {
	//boş slice koruması eklenmeli bunu unutma.
	if len(weights) == 0 {
		return 0
	}
	return calculateTotalWeight(weights) / float64(len(weights))
	//Önce calculateTotalWeight() fonksiyonunu çağır.
	//Tekrar ağırlıkları toplama.
	//Ortalamayı döndür.
}

func isEven(number int) bool {
	//Sayı çiftse true.
	return number%2 == 0
}

func countEvenNumbers(numbers []int) int {
	//isEven() fonksiyonunu kullan.
	//Çift sayıların adedini döndür.
	count := 0
	for _, number := range numbers {
		if isEven(number) {
			count++
		}
	}
	return count
}

func classifyNumber(number int) string {
	//kurallar:
	//number < 0  → "Negatif"
	//number == 0 → "Sıfır"
	//diğerleri   → "Pozitif"
	//
	//Erken return kullan.
	if number < 0 {
		return "Negatif"
	} else if number == 0 {
		return "Zero"
	} else {
		return "Pozitif"
	}

}

func isOdd(number int) bool {
	return number%2 != 0
}
func countOddNumbers(numbers []int) int {
	count := 0
	for _, num := range numbers {
		if isOdd(num) {
			count++
		}
	}
	return count
}

func isValidWeight(weight float64) bool {
	return weight > 0
}
func calculateValidWeightTotal(weights []float64) float64 {
	total := 0.0
	for _, weight := range weights {
		if isValidWeight(weight) {
			total += weight
		}
	}
	return total
}
func countValidWeights(weights []float64) int {
	count := 0
	for _, weight := range weights {
		if isValidWeight(weight) {
			count++
		}
	}
	return count
}

func calculateValidAverageWeight(weights []float64) float64 {
	//calculateValidAverageWeight
	//    ├── calculateValidWeightTotal
	//    │       └── isValidWeight
	//    │
	//    └── countValidWeights
	//            └── isValidWeight
	validCount := countValidWeights(weights)

	if validCount == 0 {
		return 0
	}

	validTotal := calculateValidWeightTotal(weights)

	return validTotal / float64(validCount)
}

func filterCalves(animals []string) []string {
	//Filtreleme fonksiyonunun amacı yalnızca buzağıları içeren yeni bir slice oluşturmaktır.
	//burayı tamamen anlamalısın.
	calves := []string{}

	for _, animal := range animals {
		if isCalf(animal) {
			calves = append(calves, animal)
		}
	}

	return calves
	//calves isimli boş sonuç listesi oluştur
	//        ↓
	//animals listesini dolaş
	//        ↓
	//hayvan buzağı mı?
	//        ↓ evet
	//sonuç listesine ekle
	//        ↓
	//döngü bitince calves listesini döndür
}

func main() {
	animals := []string{
		"İnek",
		"Buzağı",
		"Boğa",
		"Düve",
		"Buzağı",
		"İnek",
	}
	calves := filterCalves(animals)
	fmt.Println("Filtrelenen buzağılar:", calves)
	//gerçekleşen akış:
	//animals
	//    ↓
	//filterCalves
	//    ↓
	//yalnızca Buzağı değerlerinden oluşan yeni slice
	//    ↓
	//calves değişkeni

	numbers2 := []int{2, 5, 7, 8, 10, 13}

	fmt.Println("tek sayılar kaç adettir: ", countOddNumbers(numbers2))
	//örnek çıktı
	//1 - İnek
	//2 - Buzağı
	//3 - Boğa
	//4 - Düve
	//5 - Buzağı
	//6 - İnek
	//
	//Toplam hayvan: 6
	//Toplam buzağı: 2
	//Listede boğa var mı: true
	//İnek sayısı: 2
	//Toplam ağırlık: 2456.75
	//Ortalama ağırlık: 409.46
	//Çift sayı adedi: 3
	//
	//5 → Pozitif
	//12 → Pozitif
	//7 → Pozitif
	//20 → Pozitif
	//0 → Sıfır
	//-3 → Negatif

	//beklenen main akışı:
	//1. Hayvanları listele.

	printAnimals(animals)
	//2. Toplam hayvan sayısını yazdır.
	fmt.Println("Toplam hayvan: ", countAnimals(animals))
	//3. Toplam buzağı sayısını yazdır.
	fmt.Println("Toplam buzağı sayısı:", countCalves(animals))
	//4. Boğa bulunup bulunmadığını kontrol et.
	fmt.Println("Listede boğa var mı:", containsAnimal(animals, "Boğa"))
	//5. İnek sayısını yazdır.
	fmt.Println("Toplam inek sayısı", countByType(animals, "İnek"))

	weights := []float64{
		540.5,
		95.0,
		780.0,
		420.75,
		110.5,
		510.0,
	}

	//6. Toplam ağırlığı yazdır.
	fmt.Println("Toplam ağırlık: ", calculateTotalWeight(weights))
	//7. Ortalama ağırlığı yazdır.
	fmt.Println("Ortalama ağırlık: ", calculateAverageWeight(weights))

	numbers := []int{
		5,
		12,
		7,
		20,
		0,
		-3,
	}
	//8. Çift sayı adedini yazdır.
	fmt.Println("Çift sayı adedi: ", countEvenNumbers(numbers))
	//9. numbers listesindeki her sayıyı sınıflandır.
	for _, num := range numbers {
		fmt.Println("Listedeki sayıların sınıflandırılması : ", classifyNumber(num))
	}

}
