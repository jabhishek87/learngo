package main

import (
	"fmt"
)

func main() {
	fmt.Println("Everything in go pass by VALUE")
	pointer1()
	pointer2()
	pointer3()
}

func pointer1() {
	a := 43
	fmt.Println("Value of a is = ", a)
	fmt.Println("Memory address of a is = ", &a)

	var b *int = &a
	fmt.Println("Pointer refernce is = ", b)
	fmt.Println("Pointer value is = ", *b)

	// chnage value
	*b = *b + 1
	fmt.Println("new value of b is ", *b)
}
func zero(x int) {
	x = 0
}
func zero2(x *int) {
	*x = 0
}

func pointer2() {
	x := 5
	zero(x)
	fmt.Println(x)
	//fmt.Println(y)
}

func pointer3() {
	x := 5
	zero2(&x)
	fmt.Println(x)
	//fmt.Println(y)
}
