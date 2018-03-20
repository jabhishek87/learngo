package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func closure2() {
	fmt.Println("Closure2") //starts here
	x := 0                  // init
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

func closure3() {
	fmt.Println("Closure3")
	increment := wrapper()
	increment2 := wrapper()

	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(increment2())
	fmt.Println(increment2())
}

func main() {
	closure2()

	closure3()
}
