package main

import (
	"fmt"
	"strings"
)

func main() {
	var users [2]string = [2]string{
		"Khadga", "Sakar",
	}

	fmt.Println("total users:", len(users))
	fmt.Println("first user:", users[0])
	fmt.Println("last user:", users[len(users)-1])

	fmt.Println("users:", strings.Join(users[:], ","))
}
