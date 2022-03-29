package main

import (
	"fmt"
	"github.com/mjim79/go-httpclient/gohttp"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	client := gohttp.NewBuilder().
		DisableTimeouts(false).
		SetConnectionTimeout(50).
		Build()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode())
	fmt.Println(response.Status())
	fmt.Println(response.String())
}
