//operators are used to perform operation on operands
//arithmetic operator
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 39
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	//Relational Operator used to compare two operands
	var c int = 10
	var d int = 90
	fmt.Println(c == d)
	fmt.Println(c != d)
	fmt.Println(c > d)
	fmt.Println(c <= d)
	//logical operators are used to compare two or more conditions
	if a == b && c == d { //similarly || and ! are used
		fmt.Println("logically they are true")

	} else {
		fmt.Println("logically they are false")
	}
	//bitwise operator used to perform bit by bit operations
	bit := a ^ b
	fmt.Printf("%d", bit)
	//assignment operator
	c += d
	fmt.Println(c) //other assignment operators are also used

}
