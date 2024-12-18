package main

import "fmt"

func main() {
	defer fmt.Println("hiii!!") //4
	defer fmt.Println("World")  //3
	fmt.Println("Hello")        //1
	defer fmt.Println("vinit")  //2
	myDefer()

}

// defer in nutshell: execute defer code in lifo (last in first out) style.
// rest of the code is executed normally.

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
