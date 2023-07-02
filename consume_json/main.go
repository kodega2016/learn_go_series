package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	jsonData := []byte(`
	{
			"name": "khadga",
			"email": "kode@gmail.com",
			"is_active": false,
			"skills": [
					"flutter",
					"firebase",
					"react"
			]
	}`)

	isValid := json.Valid(jsonData)

	if isValid {
		var user User
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			log.Fatal("err:", err)
		}
		fmt.Println("user:", user)
		fmt.Println(user.Name, user.Email)
		fmt.Println("skills:", strings.Join(user.Skills, ","))
	} else {
		fmt.Println("json data is not valid")
	}

}

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	IsActive bool     `json:"is_active"`
	Skills   []string `json:"skills"`
}
