// constant means fixed so that it cannot be immutable further
// constant are declared using const Keyword
package main

import "fmt"
func main(){
	const a int = 100
// a := 7  value 7 can't be assigned because it is constant
	fmt.Println(a)
// There are 3 types of Constants
//numerical constant
//string literals
//boolean constant
const b float32 = 6.7
const c string = "HI squid game"
const d bool= false
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)

}