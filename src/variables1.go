package main

import "fmt"

func main()  {
  const pi float64 = 3.14
  const pi2 = 3.1415
  fmt.Println("Hello World")
  fmt.Println("Pi: ", pi)
  fmt.Println("Pi2: ",pi)

  // Declaration of interger numbers.

  base := 12
  var height int = 14
  var area int = base * height

  fmt.Println("+==============+")
  fmt.Println("Area:",area)
  
  fmt.Println("+==============+")

  // Zero values - default values

  var a int
  var b float64
  var c string
  var d bool

  fmt.Println("int:", a,
              "\nfloat64:", b, 
              "\nstring:", c, 
              "\nbool:", d)
  // Printing the types of default values
  fmt.Println("--types--")
  fmt.Printf("int a is: %T", a)
  fmt.Printf("\nfloat64 b is: %T", b)
  fmt.Printf("\nstring c is: %T", c)
  fmt.Printf("\nbool d is: %T\n", d)

  /*
    Conclusion:
      The default values never change the type
      of the declarated variable.
  */
}
