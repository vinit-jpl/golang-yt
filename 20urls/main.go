package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=grgg12423423vdfv"

func main() {
	fmt.Println("Welcome to handeling urls in golang")

	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl) // as soon as you parse the url it returns a result obj

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of qparams are %T \n", qparams) // key value pair

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=vinit",
	}

	anotherURl := partsOfUrl.String()

	fmt.Println(anotherURl)
}
