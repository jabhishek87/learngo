package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	fmt.Print("Please Enter your name: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	name := in.Text()

	fmt.Println("Hello ", name)
}
