package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to webverbs")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

// func PerformGetRequest() {
// 	const myurl = "http://localhost:3000/get"
// 	response, err := http.Get(myurl)

// 	if err != nil {
// 		panic(err)

// 	}

// 	defer response.Body.Close()
// 	fmt.Println("Status code: ", response.StatusCode)
// 	fmt.Println("Content length is:", response.ContentLength)

// 	var responseString strings.Builder
// 	content, _ := io.ReadAll(response.Body)
// 	byteCount, _ := responseString.Write(content)

// 	fmt.Println("Bytecount is: ", byteCount)
// 	fmt.Println(responseString.String())
// 	// fmt.Println(string(content))
// }

func PerformPostRequest() {
	const myurl = "http://localhost:3000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
	    	"coursename": "Lest's go with Go Lang",
			"price": 0,
			"platform": "learncodeonline.in"

		}

	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/post"

	// formdata

	data := url.Values{}
	data.Add("firstname", "vinit")
	data.Add("lastname", "jogi")
	data.Add("email", "vinit@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
