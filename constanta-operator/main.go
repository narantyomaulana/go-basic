package main

import "fmt"


func main()	{
	const NAME string = "John Doe"
	fmt.Println("Name:", NAME)	

	var nilai1 int = 30
	var nilai2 int = 20

	var hasil int = nilai1 + nilai2

	perbandingan := nilai1 > nilai2

	fmt.Println("Hasil:", hasil)
	fmt.Println("Perbandingan:", perbandingan)
}