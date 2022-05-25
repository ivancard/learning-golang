package mypackages

import "fmt"

// GuitarPublic - a guitar struct
type GuitarPublic struct {
	Brand string
	Model string
	Price int
}

//PrintMessage - print a message
func PrintMessage(text string)  {
	fmt.Println(text)
}
