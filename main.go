package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	gohttp "github.com/TestardR/go-httpclient/go-httpclient"
)

var (
	httpClient = gohttp.New()
)

func main() {

}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:last_name`
}

func getUrls() {
	headers := make(http.Header)

	headers.Set("Authorization", "Bearer ABC-123")

	response, err := httpClient.Get("https//api.github.com", headers)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Status)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := httpClient.Post("https//api.github.com", nil, user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Status)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
