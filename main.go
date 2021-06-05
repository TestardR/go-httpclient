package main

import (
	"fmt"
	"io/ioutil"

	gohttp "github.com/TestardR/go-httpclient/go-httpclient"
)

func main() {
	client := gohttp.New()

	response, err := client.Get("https//api.github.com", nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Status)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
