package main

import (
	"fmt"
	mypackages "golang-course/src/myPackages"
)

type guitar struct {
	brand string
	model string
	price int
}

func main() {
	myGuitar := guitar{brand: "Fender", model: "TeleCaster", price: 200}
	fmt.Println(myGuitar)

	// empty struct
	var secondGuitar guitar
	secondGuitar.brand = "Gibson"
	secondGuitar.model = "SG"
	secondGuitar.price = 250
	fmt.Println(secondGuitar)

	fmt.Println("===== from own package =====")
	var thirdGuitar mypackages.GuitarPublic
	thirdGuitar.Brand = "Jackson"
	thirdGuitar.Model = "Kelly"
	thirdGuitar.Price = 300
	fmt.Println(thirdGuitar)
	mypackages.PrintMessage("Hey!")
}
