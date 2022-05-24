package main

import "fmt"

func main() {
  //Declaration variables
  helloMessage := "Hello"
  wordMessage := "World"
  
  //PrintLn
  fmt.Println("<=========== PrintLn ==========>")
  fmt.Println(helloMessage, wordMessage)
  fmt.Println(helloMessage, wordMessage)

  //Printf
  fmt.Println("<=========== Printf ==========>")
  name := "Ivan"
  surname := "Cardenas"
  age := 26

  fmt.Printf("My name is %s %s and I'm %d years old\n", name, surname, age)
  fmt.Printf("My name is %v %v and I'm %v years old\n", name, surname, age)

  //Sprintf
  fmt.Println("<=========== Sprinf ==========>")
  message := fmt.Sprintf("My name is %s %s and I'm %d years old", name, surname, age)
  fmt.Println(message)

  //Type
  fmt.Printf("helloMessage is a %T type\n", helloMessage)
  fmt.Printf("age is a %T type\n", age)
} 
