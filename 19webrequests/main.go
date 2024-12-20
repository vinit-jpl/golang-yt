package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web requests in go lang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type %T", response)
	defer response.Body.Close()

	// databytes, err := ioutil.ReadAll(response.Body) // ioutil is depricated
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
