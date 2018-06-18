package main

import (
	"fmt"
)

// https://stackoverflow.com/questions/41668053/cap-vs-len-of-slice-in-golang

func main() {
	go_map := make(map[string]int)
	go_map["first_key"] = 25

	go_map["2nd_key"] = 30

	// remove key
	delete(go_map, "first_key")

	if element, ok := go_map["first_key"]; ok {
		fmt.Println(element) // won't be printed
	}

	if element, ok := go_map["2nd_key"]; ok {
		fmt.Println(element)
	}

}