package main

import "fmt"

func main() {
	kode := Person{
		name:     "khadga bahadur shrestha",
		email:    "khadgalovecoding2016@gmail.com",
		isActive: true,
	}

	fmt.Println("kode:", kode)
	fmt.Println(kode.name, kode.email)

	kode.updateName("Arjun Shrestha")
	kode.print()
}

type Person struct {
	name     string
	email    string
	isActive bool
}

func (p *Person) updateName(name string) {
	p.name = name
}

func (p Person) print() {
	fmt.Println(p.name, p.email)
}
