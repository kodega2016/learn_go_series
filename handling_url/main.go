package main

import (
	"fmt"
	"log"
	"net/url"
)

const myurl = "https://jsonplaceholder.typicode.com/users/1"

func main() {
	parsed, err := url.Parse(myurl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hostname:", parsed.Host)
	fmt.Println("path:", parsed.Path)
	fmt.Println("query:", parsed.RawQuery)
	fmt.Println("schema:", parsed.Scheme)

	myweb := &url.URL{
		Scheme:   "https",
		Host:     "google.com",
		Path:     "/flutter",
		RawQuery: "query=state-management",
	}

	fmt.Println(myweb)
}
