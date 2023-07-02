package main

import "fmt"

func main() {
	var p *int
	x := 10
	p = &x

	fmt.Println("address:", p, "value:", *p)
}
