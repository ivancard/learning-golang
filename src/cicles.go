package main

import "fmt"

func main() {

	//for conditional
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	//for while
	var counter int
	for counter <= 10 {
		fmt.Printf("%d ", counter)
		counter++
	}
	fmt.Println()

	//for forever
	var conterForever int
	for {
		fmt.Println(conterForever)
		conterForever++
	}
}
