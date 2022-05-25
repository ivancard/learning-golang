package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")

}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPc pc) Stringify() string {
	return fmt.Sprintf("The pc have \t %d GB of RAM\n\t\t %d TB of disk\n\t\t %s is your brand", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	a := 50
	b := &a

	fmt.Println("a is: ", a)                                    //50
	fmt.Println("b is a pointer to a, and the position is:", b) //hexcode
	fmt.Printf("b is a type: %T\n", b)                          //*int
	fmt.Println("the content of pointer b is:", *b)             //50

	*b = 51
	fmt.Println("the content of a modified by his pointer is:", a)
	fmt.Println("=====================")
	myPc := pc{ram: 16, disk: 5, brand: "IBM"}
	// fmt.Println(myPc)
	myPc.ping()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)

	// Print with a formater function
	fmt.Println("=====================")
	fmt.Println(myPc.Stringify())
}
