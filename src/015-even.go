package main

import (
	"fmt"
)

func main() {
	fmt.Print("Please Enter a UpperLimit: ")
	var upperLimit int
	fmt.Scanf("%d", &upperLimit)

	// Logic 1 with no condition
	for i:=2; i<=upperLimit; i+=2 {
		fmt.Print(i, " ")
	}

	fmt.Println(" ")
	// logic two with modulus
	for i:=1; i<=upperLimit; i++ {
		if i%2 == 0{
			fmt.Print(i, " ")
		}
	}
}
