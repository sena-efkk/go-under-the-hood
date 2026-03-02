package main

import "fmt"

/*

Senaryo:

    OrderStatus adında, temeli int olan bir tip tanımla.

    const bloğu içinde iota kullanarak şu durumları
	tanımla: Pending, Processing, Shipped, Delivered.

    Bir fonksiyon yaz (veya main içinde yap): Pending durumundaki
	bir siparişin int değerini ve durum ismini ekrana yazdır.

    Kritik Soru: Eğer Shipped durumuna 1 eklemeye çalışırsan (çünkü arkada int),
	doğrudan ekleyebilir misin yoksa OrderStatus(1) şeklinde mi eklemelisin? Dene ve gör.


*/

//iota, const blokları içinde kullanılan ve her satırda otomatik olarak 1 artan bir sayacı temsil eder.
const (
	Pending OrderStatus = iota // 0
	// İlk sabit tipini belirler, sonrakiler onu takip eder
	Processing // 1
	Shipped    // 2
	Delivered  //3
)

type OrderStatus int

func main() {

	fmt.Println(int(Pending))
	//doğru
	fmt.Println(Shipped + 1)
	//doğru.
	fmt.Println(OrderStatus(Shipped) + 1)
	//Doğru. Tip dönüşümü yapsan da matematiksel değer
}
