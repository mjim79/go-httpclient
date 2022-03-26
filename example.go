package main

import (
	"fmt"
	"github.com/mjim79/go-httpclient/gohttp"
	"io/ioutil"
	"net/http"
)

var httpClient = gohttp.New()

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := httpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {

	response, err := httpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
