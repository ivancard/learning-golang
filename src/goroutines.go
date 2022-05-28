package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {

	var wg sync.WaitGroup

	// fmt.Println(wg)

	fmt.Println("Hello")
	wg.Add(1)
	// fmt.Println(wg)
	go say("World", &wg)

	wg.Wait()

}
