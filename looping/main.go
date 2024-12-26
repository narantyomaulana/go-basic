package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	// While loop
	i := 0
	for i < 5 {
		fmt.Println("Perulangan2 ke-", i)
		i++
	}

	// Infinite loop
	j := 0
	for {
		fmt.Println("Perulangan3 ke-", j)
		j++
		if j == 5 {
			break
		}
	}
}