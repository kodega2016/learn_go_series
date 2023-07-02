package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	message := "Thank you for visting,please leave rating"
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Thanks for rating...", numRating)
}
