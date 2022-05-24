package main

import "fmt"

/*
  Some Keyword of golang.
  defer, break, continue
*/

func main() {
	//Defer
	defer fmt.Println("Hello")
	fmt.Println("World")

	//Continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Is five")
			continue
		}
		if i == 8 {
			fmt.Println("Break")
			break
		}
		fmt.Println(i)
	}

}
