package main

import "fmt"
import "strings"

func isPalindrome(text string)bool {
  text = strings.ToLower(text)
  var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i]) 
	}
	if textReverse == text {
	  return true
	}
	return false
}

func main() {
	//Array
	fmt.Println("========== Arrays ==========")
	var arr [4]int
	arr[0] = 5
	arr[1] = 5
	fmt.Println(arr, len(arr), cap(arr))

	//Slice
	fmt.Println("========== Slice ==========")
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Slice selection
	fmt.Println("========== Splice selection ==========")
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Slice methods
	fmt.Println("========== Splice methods ==========")
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append another slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//for range
	fmt.Println("========== for range ==========")
	for _, value := range slice {
		fmt.Println(value)
	}

	//isPalindrome function
	fmt.Println(isPalindrome("Ana"))
}
