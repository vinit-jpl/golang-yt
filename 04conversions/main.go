package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app!")
	fmt.Println("Please rate our pizza between 1-5.")

	reader := bufio.NewReader(os.Stdin) // stores the referenc of input

	input, _ := reader.ReadString('\n')

	// input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("Thanks for rating,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating:", numRating+1)
	}
}
