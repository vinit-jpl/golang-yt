package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	vinit := User{"Vinit", "vinit@go.dev", true, 23}
	fmt.Println(vinit)
	fmt.Printf("Vinit details are: %+v\n", vinit) //%+v gives more detailed struct
	fmt.Printf("Name is %v and email is %v \n", vinit.Name, vinit.Email)
	vinit.GetStatus()
	vinit.newMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) newMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
