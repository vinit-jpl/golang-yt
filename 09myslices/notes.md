## Slices

var fruitList = []string{"apple", "tomato", "peach"}
fmt.Printf("the type of fruitList is %T\n", fruitList)

fruitList = append(fruitList, "mango", "banana")
fmt.Println(fruitList) ===> output: [apple tomato peach mango banana]

fruitList = append(fruitList[1:])  ===> start from index 1 and go till the end of the list.
fmt.Println(fruitList)    ===> output: [tomato peach mango banana]



## syntax 
variable = append("list"[start_index, end_index])

end_index is non inclusive.

## memory allocation using make();

syntax: make([]<data-type>, size)

## Appending Beyond the Original Length go

highScores = append(highScores, 555, 666, 777)

The append function is special. When you append elements to a slice, Go checks if the underlying array has enough capacity to hold the new elements:
If the capacity is sufficient, the new elements are added, and the slice's length is updated.
If the capacity is insufficient, a new, larger array is allocated. The data from the original array is copied into the new array, and the new elements are added. This process is known as memory reallocation.


## removing a value from a slice based on index
