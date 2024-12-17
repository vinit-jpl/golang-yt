package main

import "fmt"

func main() { // main function is entry point of fucntion.
	fmt.Println("Welcome to fucntions in goLang!!")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
}

func greeter() {
	fmt.Println("Hello from go lang")
}

func adder(valOne int, valTow int) int {
	return valOne + valTow
}

// case : Dont know how many values come in but want to add all values

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}
