package main

import "fmt"

type guitar struct {
	brand string
	model string
	price int
}

func main() {
	myGuitar := guitar{brand: "Fender", model: "TeleCaster", price: 200}
	fmt.Println(myGuitar)
}
