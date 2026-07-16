package main

import "fmt"

//func calculate(a int, b int) (int, int) {
//	return a + b, a * b
//}
//func divideWithRemainder(a int, b int) (int, int) {
//	return a / b, a % b
//}

//func add(a, b int) int :ekle
//func subtract(a, b int) int :çıkar
//func multiply(a, b int) int :çarp
//func calculateRemainingCapacity(maxCapacity, currentCount int) int :kalankapasiteyi hesapla

//	func add(a, b int) int {
//		return a + b
//	}
//
//	func substract(a, b int) int {
//		if a > b {
//			return a - b
//		} else {
//			return b - a
//		}
//	}
//
//	func multiply(a, b int) int {
//		return a * b
//	}
//
//	func calculateRemainingCapacity(maxCapacity, currentCount int) int {
//		return maxCapacity - currentCount
//	}
//
//	func classifyWeight(weight float64) string {
//		//ağırlık sınıflandırılması:
//		//weight <= 0   → "Geçersiz"
//		//weight < 100  → "Hafif"
//		//weight < 400  → "Orta"
//		//diğerleri     → "Ağır"olucak
//		if weight <= 0 {
//			return "Geçersiz"
//		} else if weight < 100 {
//			return "Hafif"
//		} else if weight < 400 {
//			return "Orta"
//		} else {
//			return "Ağır"
//		}
//	}
func classifyWeight(weight float64) string {
	if weight <= 0 {
		return "Geçersiz"
	}

	if weight < 100 {
		return "Hafif"
	}

	if weight < 400 {
		return "Orta"
	}

	return "Ağır"

	//bu kod diğerinden farkı erken return olayını kullanmasıdır yani else yok
	//Fonksiyon normal sonuna ulaşmadan, gerekli cevap bulunduğu anda döner.
}

//func isEven(number int) bool {
//	return number%2 == 0
//}
//func countEvenNumbers(numbers []int) int {
//	count := 0
//	for _, number := range numbers {
//		if isEven(number) {
//			count++
//		}
//	}
//	return count
//}

func calculateTotalWeight(weights []float64) float64 {
	total := 0.0
	for _, weight := range weights {
		total += weight
	}
	return total
}
func calculateAverageWeight(weights []float64) float64 {
	//bu if olayına kısa koruma yani,
	if len(weights) == 0 {
		return 0
	}

	return calculateTotalWeight(weights) / float64(len(weights))
}

func main() {

	agirlikdizim := []float64{5.01, 6.5, 4.0, 8}
	fmt.Println(calculateTotalWeight(agirlikdizim))
	fmt.Println(calculateAverageWeight(agirlikdizim))

	//
	//ciftsayilar := []int{
	//	4, 23, 56, 2, 99, 7, 2, 2, 2, 2,
	//}
	//fmt.Println(countEvenNumbers(ciftsayilar))
	//fmt.Println(isEven(ciftsayilar[1]))

	//Boş slice nedir?
	//
	//Boş slice, içinde hiç eleman olmayan listedir.

	//fmt.Println(classifyWeight(7400.7))

	//fmt.Println(calculateRemainingCapacity(500, 3)) // burada hatan var
	//fmt.Println(add(5, 5))
	//fmt.Println(substract(99, 1))
	//fmt.Println(multiply(6, 6))

	//fmt.Println(calculate(1, 2))
	//sum, multiplication := calculate(4, 4)
	//fmt.Println("toplam: ", sum)cd
	//fmt.Println("çarpım: ", multiplication)
	//
	//quotient, remainder := divideWithRemainder(17, 5)
	//fmt.Println("Bölüm:", quotient)
	//fmt.Println("Kalan:", remainder)

}
