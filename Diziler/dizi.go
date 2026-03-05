package main

import "fmt"

func main() {
	//array
	var dizi [3]int = [3]int{10, 20, 30}
	dizi[0] = 100
	fmt.Println(dizi[0])
	fmt.Println(dizi)

	students := [4]string{"sena", "erva", "furkan", "timur"}
	students[0] = "sultan"
	fmt.Println(students)

	//slice:eleman ekleme çıkarma yapabilirim.
	salaries := []int{100, 200, 150, 500}
	salaries = append(salaries, 1000)
	fmt.Println(salaries)
	fmt.Println(salaries[0:2]) //bana 0. indexten 2. indexe kadar olanları getir 2.index dahil değil
	//salaries = salaries[0:2]   //yeni dizimiz 100 , 200 olur
	fmt.Println(len(salaries))

}
