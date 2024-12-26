package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	student1 := student{
		name: "John Doe",
		age:  25,
	}

	student2 := student{}
	student2.name = "Jane Doe"
	student2.age = 24

	fmt.Println("Name :", student1.name)
	fmt.Println("Age :", student1.age)

	fmt.Println("=====================================")

	fmt.Println("Name :", student2.name)
	fmt.Println("Age :", student2.age)

	students := []student{
		{name: "Akbar", age: 26},
		{name: "Dimas", age: 27},
		{name: "Rizky", age: 28},
	}

	fmt.Println("=====================================")

	for _, s := range students {
		fmt.Println("Name:", s.name)
		fmt.Println("Age:", s.age)
		fmt.Println("=====================================")
	}

}
