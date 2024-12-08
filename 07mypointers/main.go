package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers!")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("value of actual pointer is", ptr)
	fmt.Println("value of actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value", myNumber)
	fmt.Println("Address", &ptr, &myNumber)

}
