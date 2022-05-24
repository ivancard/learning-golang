package main

import "fmt"

func main() {
 
  x := 10
  y := 50

  //suma
  result := x+y
  fmt.Println("Suma:", result)

  //Resta
  result = y-x
  fmt.Println("Resta:", result)

  //Multiplicacion
  result = y * x
  fmt.Println("Multiplicacion:", result)

  //Division
  result = y / x
  fmt.Println("Division:", result)

  //Modulo
  result = y % x
  fmt.Println("Modulo:", result)

  //Incremental 
  x++
  fmt.Println("Incremental:", x)

  //Decremental
  y--
  fmt.Println("Decremental:", y)
  

}
