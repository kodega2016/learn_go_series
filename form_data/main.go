package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	const myurl = "https://jsonplaceholder.typicode.com/todos"

	data := url.Values{}
	data.Add("title", "Flutter and Golang")
	data.Add("completed", "false")
	data.Add("userId", "2")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		log.Fatal(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	responseData := string(bs)

	fmt.Println("response data:", responseData)

}
