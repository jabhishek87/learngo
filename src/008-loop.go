package main

import "fmt"

func main() {
	for i:=0; i<=100; i+=2 {
		fmt.Println(i)
	}

	// Range loop

	names := []string{
		"Abhishek",
		"Priti",
		"Dhrishti",
	}

	for i, name  := range names {
		fmt.Printf("%d, %s\n", i, name)
	}
}
