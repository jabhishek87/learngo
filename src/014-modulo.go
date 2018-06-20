package main

import (
	"fmt"
)

func main() {
	fmt.Print("Please Enter any Two number: ")
	var num1, num2 int
	num2 = 1 // default value
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Println(num1, num2)
	fmt.Println(num1, "%", num2, "=", num1 % num2)
}
