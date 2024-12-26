package main

import "fmt"

func main() {
	// Array
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Mango"
	fruits[3] = "Pineapple"
	fruits[4] = "Grape"

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Array Slice

	var cars = []string{"BMW", "Mercedes", "Audi", "Toyota", "Honda"}

	for _, car := range cars {
		fmt.Println(car)
	}

	var cars2 = cars[1:3] // index 0 to 1
	fmt.Println(cars2)

}
