package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Objective:
// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	fmt.Print("Enter address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading address:", err)
		return
	}

	data := map[string]string{
		"name":    strings.TrimSpace(name),
		"address": strings.TrimSpace(address),
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
