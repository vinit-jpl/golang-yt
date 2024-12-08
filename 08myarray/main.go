package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to array in Go")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("length of Fruit list:", len(fruitList)) // length depends on size declared

	var vegList = [5]string{"potato", "beans", "onion"}
	fmt.Println(vegList, len(vegList))

}
