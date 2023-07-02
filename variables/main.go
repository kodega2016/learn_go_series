package main

import "fmt"

const PI = 3.14

func main() {
	var name string = "Let's go with Golang"
	fmt.Println(name)

	var author string = "Hitesh Chadhary"
	fmt.Println(author)

	channelName, isPaid := "Hitesh Chaudhary", false
	fmt.Println(channelName, isPaid)

	var (
		bookName   string = "Explore Golang"
		bookAuthor string = "Hitesh Chaudhary"
	)

	fmt.Println(bookName, bookAuthor)

	fmt.Println("The value of PI:", PI)
}
