package main

import "fmt"

func main() {
	a := "gopi"
	b := "gopi"
	if a == b {
		fmt.Println("They are same") //using only if
	}
	c := 4
	d := 10
	if c == d {
		fmt.Println("They are same")

	} else {
		fmt.Println("They are different") //using if and else
	}
	if c > d {
		fmt.Println("a is greater")
	} else if c == d {
		fmt.Println("They are equal") //using if else and if_else
	} else {
		fmt.Println("nothing true")
	}
}
