package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welocme to files in go lang")
	content := "this needs to be go in a file"

	file, err := os.Create("./myLcoGoFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("lenght is: ", length)
	defer file.Close()
	readFile("./myLcoGoFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("text data inside the file is\n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err) // panic stops the execution of the program and throws error
	}
}
