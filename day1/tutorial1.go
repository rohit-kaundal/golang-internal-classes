package main

import "fmt"

func FirstDay() {
	var fname string
	sname := "Rohit"
	sname = "Monika"
	fname = "Deepak"
	age1, age2 := 30, "32"

	age3 := float32(age1) + 0.2
	fmt.Println(fname, sname, age1, age2, age3)

	if age1 > 20 {
		fmt.Println("Age is greater then 20")
	}

	if age5 := age3 + 0.3; age5 > 30 {
		fmt.Println("Age is greater then 30")
	}

	if fname2, lname2, err := FullName(); err != nil {
		fmt.Println("Error occured")
		fmt.Println(fname2, lname2)
	}

	f, s, _ := FullName()
	fmt.Println(f, s)

	for i := 0; i < 10; i++ {

		fmt.Println(i)
		if i%2 == 0 {
			break
		}
	}

	j := 0
	for j < 10 {
		j++
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

}

func FullName() (string, string, error) {
	return "rohit", "kaundal", nil
}
