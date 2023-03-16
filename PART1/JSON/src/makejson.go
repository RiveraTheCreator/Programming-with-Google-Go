package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Ask the user to enter their name and address
	var name, address string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your address: ")
	fmt.Scanln(&address)

	// Create a map with the entered data
	data := map[string]string{
		"name":  name,
		"address": address,
	}

	// Create a JSON object from the map using the Marshall() function
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error creating JSON object:", err)
		return
	}

	// Print the JSON object to console
	fmt.Println(string(jsonData))
}
