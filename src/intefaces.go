package main

import "fmt"

type figures2d interface {
  area() float64
}

type square struct {
	base float64
}
type rectangle struct {
	width  float64
	height float64
}

func (sq square) area() float64 {
	return sq.base * sq.base
}
func (rec rectangle) area() float64 {
	return rec.width * rec.width
}
func calculate(f figures2d)  {
  fmt.Println("Area:", f.area())
}

func main() {
	mySquare := square{base: 20.5}
	myRectangle := rectangle{width: 15, height: 10}

	calculate(mySquare)
	calculate(myRectangle)
	fmt.Println("===========")
	myIterface := []interface{}{"hey!", 50, true}
	fmt.Println(myIterface...)
}
