package main

import "fmt"

//iota, const blokları içinde kullanılan ve her satırda otomatik olarak 1 artan bir sayacı temsil eder.
const (
	Beklemede  = iota // 0
	Onaylandi         // 1
	Reddedildi        // 2
)

func main() {

	fmt.Println(Beklemede, Onaylandi, Reddedildi)
}
