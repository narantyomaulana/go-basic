package main

import "fmt"

func main() {
	value := 400
	value2 := 200

	if value > value2 {
		fmt.Println("Value is greater than value 2")
	} else {
		fmt.Println("Value is less than value 2")
	}

	const address string = "Bandung, Indonesia"

	switch address {
	case "Jakarta, Indonesia":
		fmt.Println("Address is in Jakarta, Indonesia")
	case "Bandung, Indonesia":
		fmt.Println("Address is in Bandung, Indonesia")
	default:
		fmt.Println("Address is unknown")
	}

}
