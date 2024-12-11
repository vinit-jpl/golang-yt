package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	vinit := User{"Vinit", "vinit@go.dev", true, 23}
	fmt.Println(vinit)
	fmt.Printf("Vinit details are: %+v\n", vinit) //%+v gives more detailed struct
	fmt.Printf("Name is %v and email is %v: \n", vinit.Name, vinit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
