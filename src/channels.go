package main

import "fmt"

func sayAgain(text string, c chan<- string) {
	c <- text
}
func printChannel(c <-chan string)  {
  fmt.Println(<-c)
}
func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go sayAgain("Bye", c)

	// fmt.Println(<-c)
	printChannel(c)
}
