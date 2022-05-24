package main

import (
	"fmt"
	"strconv"
)

func main() {
  value, err := strconv.Atoi("4u5")
  if err != nil{
    fmt.Println(err)
    return
  }
  
  fmt.Println("Value:", value)
}
