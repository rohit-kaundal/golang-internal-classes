package main

import (
	"fmt"
)

func Run() {
	fmt.Println("Working Fine")

	// Array - fixed length
	var marks [5]int
	marks[0] = 100
	marks[1] = 200
	fmt.Println(marks)

	// Slice - not fixed length
	var ages []int
	ages = append(ages, 30, 40, 50, 60, 70)
	ages[1] = 50
	fmt.Println(ages)

	// array inline definition
	var years = [3]int{1991, 1992, 1993}
	years[1] = 2023
	fmt.Println(years)

	// slice inline definition
	var names = []string{"Rohit", "Deepak", "Monika"}
	names = append(names, "Rohit Kaundal")

	for _, name := range names {
		fmt.Println(name)
	}

	// HashMap : e.g age[0] = 34, name[0] = "Rohit Kaundal", data["RohitKaundal"] = 34, data["RohitKaundal"]["Age"] = 34

	var data1 = map[string]int{}

	data1["RohitKaundal"] = 34
	data1["Gaurav"] = 30
	fmt.Println(data1)

	for key, val := range data1 {
		fmt.Println("Key is:", key, "and value is", val)
	}

	// Hashmap with multipe keys
	data2 := make(map[string]map[string]int)

	data2["rohitkaundal"] = make(map[string]int)
	data2["rohitkaundal"]["age"] = 34
	data3 := data2["rohitkaundal"]["age"]
	fmt.Println(data3)
}
