package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")
	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("the type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // start from value at index 1 and go till index 2.
	fmt.Println(fruitList)

	highScores := make([]int, 4) // default allocation of memory
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 // will throw error index out of bound

	highScores = append(highScores, 555, 666, 777) // does not throw error coz when we use appedn rellocation of memeory takes place
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
