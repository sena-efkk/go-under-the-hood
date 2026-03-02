package main

import "fmt"

type Temperature float64
type DeviceID uint16
type DeviceStatus uint8

// deneme
const (
	Offline     DeviceStatus = iota // 0
	Online                          // 1
	Maintenance                     // 2
	Faulty                          // 3
)

func main() {
	// 1. Ham veriyi kendi tipimize sokuyoruz
	var rawData uint8 = 100
	temp := Temperature(rawData) + 0.5

	// 2. DeviceID tanımlama
	var id DeviceID = 1024

	// 3. Status ve Sayı Toplama (Kritik Nokta)
	// Faulty bir DeviceStatus'tür (uint8). 1 ise bir int sabiti.
	// Go sabitlerde (constants) esnektir, bu yüzden direkt toplayabilirsin.
	newStatus := Faulty + 1

	// 4. String Birleştirme (Mühendislik Çözümü)
	// Go'da "+" ile farklı tipleri birleştiremezsin.
	// Ya fmt.Println'e virgülle verirsin ya da fmt.Sprintf kullanırsın.
	bilgi := fmt.Sprintf("Cihaz ID: %d | Sıcaklık: %.1f | Yeni Durum Kodu: %d", id, temp, newStatus)

	fmt.Println(bilgi)

}
