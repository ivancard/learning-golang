package main

import "fmt"

func main() {
	switch 	module := 5 % 2; module {
	case 0:
		fmt.Println("Is even")
	default:
		fmt.Println("Is odd")
	}

	//without condition
	fmt.Println("<======whitout condition=====>")
	value := 50
	switch  {
	case value > 50:
	  fmt.Println("Is upper to 50")
	case value < 50:
	  fmt.Println("Is lower to 50")
	default:
	  fmt.Println("Is 50")
	 
	}
	
}
