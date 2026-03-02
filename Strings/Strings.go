package main

import (
	"fmt"
	"unicode/utf8"
)

//strings ve runes.
//rune, aslında int32 tipinin bir takma adıdır (alias).
//Bir rune, tek bir Unicode karakterini temsil eder.
//Örneğin 'A' bir rune'dur, '€' bir rune'dur, 'ğ' bir rune'dur.

func main() {
	metin := "Golang Öğreniyorum"
	fmt.Println("byte uzunlugu sayısı: ", len(metin))
	fmt.Println("rune sayısı(gerçek karakter sayısı): ", utf8.RuneCountInString(metin))

}
