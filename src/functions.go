package main

import "fmt"

func printMessage( message string){
  fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
  fmt.Println(a, b, c) 
}

func returnSquare(num int)int  {
 return num * num 
}

func doubleReturn(num int)(c, d int){
 return num, num * num
}

func main() {
  printMessage("Hello World")
  tripleArgument(1, 2, "are numbers")
  fmt.Println(returnSquare(2))
  fmt.Println("=============")

  initial, square := doubleReturn(4)
  fmt.Printf("the number %d has a square of %d\n", initial, square)

  _ , square2 := doubleReturn(4)
  fmt.Printf("the square of a number is: %d\n", square2)

}
