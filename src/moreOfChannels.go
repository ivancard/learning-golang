package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"

	//Range and close

	close(c)

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mail 1", email1)
	go message("mail 2", email2)

	for i := 0; i < 2; i++ {
	 select {
	 case e1 := <-email1:
	   fmt.Println("Email recibido de email1:", e1)
	 case e2 := <-email2:
	   fmt.Println("Email recibido de email1:", e2)
	 }
	}

}
