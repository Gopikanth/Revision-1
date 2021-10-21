package main

import "fmt"

func main() {
	i := 0
	for i < 5 { //loop runs untill condition become false
		i ++
	}
	fmt.Println(i)


//for in the range
 a := [5]int{ 1,2,3,4,5}
for ele , ok := range a {
	fmt.Println(ele,ok)

}
//for in mapping
 maping := map[int]string{
        22:"Gopi",
        33:"kanth",
        44:"BE",
    }
	for ele , ok := range maping {
	fmt.Println(ele,ok)
	}
//for in channel
channel := make(chan int)
    go func(){
        channel <- 10
        channel <- 20
       channel <- 30
       channel <- 40
       close(channel)
    }()
    for i:= range channel {
       fmt.Println(i) 
    }
}
