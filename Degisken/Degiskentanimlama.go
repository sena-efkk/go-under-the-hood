package main

import "fmt"

//ekrana yazı yazdırmak için I/O işlemleri için sart

//furc main programın giriş noktasıdır .uyg burdan başlar.
func main() {

	sayi := 10     //int
	ondalik := 5.5 // float64

	sonuc := float64(sayi) + ondalik
	fmt.Println(sonuc)

	//---------------------------------

	// 1. uint8 tipinde 255 değerinde bir değişken tanımla
	var sensorVerisi uint8 = 255

	// 2. Bunu float64'e dönüştür (Dönüşüm syntax: float64(degisken))
	floatVeri := float64(sensorVerisi)

	// 3. Üzerine 10.5 ekle
	sonuc2 := floatVeri + 10.5
	fmt.Println("Sonuç:", sonuc2)

	// 4. BONUS: uint8 olan sensorVerisi'ne 1 ekle ve ekrana yazdır
	// Bakalım 255 + 1 ne yapacak? (Taşma/Overflow kavramı)
	sensorVerisi = sensorVerisi + 1
	fmt.Println("Bonus Sonuç:", sensorVerisi)
}
