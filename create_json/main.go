package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	users := []User{
		{"khadga", "kode@gmail.com", "passcode", false, []string{"flutter", "firebase", "react"}},
		{"sakar", "rakasidebus@gmail.com", "passcode", false, []string{"node js", "react js", "css"}},
	}

	// res, err := json.Marshal(users)

	res, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", res)

}

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	IsActive bool     `json:"is_active"`
	Skills   []string `json:"skills"`
}
