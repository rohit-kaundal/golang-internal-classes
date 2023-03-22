package main

import "fmt"

// User defined types Struct Interfaces
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var ptrAge *int

func (p *Person) IncermentAge(val int) {
	p.Age += val
}

func (p Person) IsAdult() bool {
	return p.Age > 18
}

func main() {
	fmt.Println("Day 3")

	monika := Person{FirstName: "Monika", LastName: "Babutta", Age: 32}

	fmt.Println(monika)

	rohitkaundal := Person{}
	rohitkaundal.FirstName = "Rohit"
	rohitkaundal.LastName = "Kaundal"
	rohitkaundal.Age = 34
	rohitkaundal.IncermentAge(1)

	fmt.Println(rohitkaundal)
	fmt.Println("Is Adult", rohitkaundal.IsAdult())

	// Pointers

	year := 50
	ptrYear := &year

	fmt.Println(year, *ptrYear)

	// Switch example

	switch {
	case monika.Age < rohitkaundal.Age:
		fmt.Println("Younger then Rohit K")
	case monika.Age < 18:
		fmt.Println("Minor")
	}

	switch i := monika.Age; i {
	case 32:
		fmt.Println("her age is 34")
	case 18:
		fmt.Println("is adult")
	}

	fmt.Println(addMultipleIntegers(1, 2))
	fmt.Println(addMultipleIntegers(1, 2, 3, 4))

}

// variadic functions

func addMultipleIntegers(n ...int) int {
	sum := 0
	for _, val := range n {
		sum += val
	}
	return sum
}
