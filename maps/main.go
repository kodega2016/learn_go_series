package main

import "fmt"

func main() {
	var userInfo map[string]string
	userInfo = make(map[string]string)

	// assigning value to the map
	userInfo["name"] = "khadga bahadur shrestha"
	userInfo["role"] = "software developer"
	userInfo["address"] = "madhumalla morang"
	userInfo["password"] = "password"

	// deleting value from the map
	delete(userInfo, "password")

	// get the value from the map
	name, isOk := userInfo["name"]
	if isOk {
		fmt.Println("name:", name)
	}

	// iterate over the map
	for key, value := range userInfo {
		fmt.Println(key, value)
	}

	fmt.Println("total key-value pair:", len(userInfo))

	fmt.Println(userInfo)

}
