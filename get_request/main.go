package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	res, err := http.Get(myurl)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bs, e := ioutil.ReadAll(res.Body)

	if e != nil {
		log.Fatal(e)
	}

	response := string(bs)
	fmt.Println("response:", response)
}
