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

}
