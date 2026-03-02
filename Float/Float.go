package main

import "fmt"

// 1. Yeni tip tanımı (Fonksiyonun dışında veya içinde olabilir)
type Money float64

//float64 sadece bir sayıdır. Ama Money bir kavram olmuştur.
func main() {
	var maas Money = 5000.0
	var bonus float64 = 250.75

	// Bu satır hata verecek:
	// toplam := maas + bonus

	// Çözüm: Açıkça dönüştür
	toplam := maas + Money(bonus)

	fmt.Println("Toplam:", toplam)
}
