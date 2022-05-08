package main

import "fmt"

func maps() {

	var user map[string]string = map[string]string{
		"name":     "Sanya",
		"lastName": "Seredavkin",
	}

	sizedMap := make(map[int]string, 10)

	userCount := len(user)

	//when key doesn't exist will return default value

	fName := user["firstName"] // will return empty string

	// how to check key existing

	fName, fNameExist := user["firstName"]

	if !fNameExist {
		fmt.Println("firstName doesn't exists")
	}

	//remove key

	delete(user, "firstName")

	fmt.Println(sizedMap, userCount, fName)
}
