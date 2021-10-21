// some main Data Types are
//Number
//Boolean
//string
//float
//commplex
package main

import "fmt"

func main() {
	a := "This is a string"
	var b int = 60
	var c bool = true
	var d float64
	d = 5.5
	var e complex128 = complex(5, 5)
	fmt.Println(a)
	fmt.Printf("This is of type %T\n", a)
	fmt.Println(b)
	fmt.Printf("This is of type %T\n", b)
	fmt.Println(c)
	fmt.Printf("This is of type %T\n", c)
	fmt.Println(d)
	fmt.Printf("This is of type %T\n", d)
	fmt.Println(e)
	fmt.Printf("This is of type %T\n", e)
}
