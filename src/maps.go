package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Carl"] = 23
	m["Joseph"] = 25
	fmt.Println(m)

	//Iterate map
	fmt.Println("==================")
	for index, value := range m {
		fmt.Println(index, value)
	}

	//Print one value
	fmt.Println("==================")
	value := m["Carl"]
	value2, ok := m["Cardl"]
	fmt.Println(value)
	fmt.Println(value2,ok)
}
