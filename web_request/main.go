package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		log.Fatal("error:", err)
	}

	defer res.Body.Close()

	bs, _ := ioutil.ReadAll(res.Body)
	value := string(bs)

	fmt.Println("response:", value)
}
