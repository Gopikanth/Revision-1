// variables are the containers to store the values
//In GO variables should start with Letter or Under_score
//variables are created using Var keyword

package main

import "fmt"

func main() {
	var binu int = 60  //but 1binu is illegal
	var _c bool = true // #c is illegal
	A := 10
	a := 11
	
	fmt.Println(binu)
	fmt.Println(_c)
	fmt.Println(A)
	fmt.Println(a) //variables are also case sensitive
	//a is not same as A
    gopi , kanth := 2, 5  //multiple variables can be declared in same line in this short form
	fmt.Println(gopi)
	fmt.Println(kanth)
	kloud , one := 100 , "percentage"// different dat types can also be declared in this manner
	fmt.Println(kloud)
	fmt.Println(one)

}

