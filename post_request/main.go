package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos"

	rawData := strings.NewReader(`
	{
		"userId": 1,
		"id": 1,
		"title": "delectus aut autem",
		"completed": false
		}
	`)

	res, err := http.Post(myUrl, "application/json", rawData)

	if err != nil {
		log.Fatal(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	bsString := string(bs)

	fmt.Println("response:", bsString)
}
