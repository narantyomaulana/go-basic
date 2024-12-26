package main

import "fmt"

func main() {

	// Declare a variable
	var number1 int = 10
	
	fmt.Println("Value of number1:", number1)
	fmt.Println("Address of number1:", &number1)

	var number2 *int = &number1

	fmt.Println("Value of number2:", number2)
	fmt.Println("Address of number2:", *number2)

	fmt.Println("=====================================")

	// Change the value of number1
	number1 = 20
	fmt.Println("Value of number1:", number1)
	fmt.Println("Address of number1:", &number1) // Address of number1 is still the same
	fmt.Println("Value of number2:", number2)
	fmt.Println("Address of number2:", *number2) // Value of number2 is still the same
}