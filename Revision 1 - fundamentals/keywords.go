package main

import "fmt"

func main() {
	var a string = "Hi This is again identifier"
	fmt.Println(a)
}

//keywords are the reserved words that are used only for internal process so that it cannot be used as identifier
// for example if we write var defer string = "Hi This is correct" 
//it will throw error because defer is a keyword
