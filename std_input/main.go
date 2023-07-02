package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := "Enter your full name"
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("welcome:", name)

}
