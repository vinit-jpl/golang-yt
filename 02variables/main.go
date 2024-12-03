package main

import "fmt"

// token := 30000 can't use outside any method

func main() {
	var user string = "Vinit"
	fmt.Println(user)
	fmt.Printf("The varibale is of Type %T \n", user)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The varibale is of Type %T \n", isLoggedIn)

	// implicit type

	var website = "google.com"
	fmt.Println(website)

	// no var style

	name := "vinit"
	fmt.Println(name)

}
