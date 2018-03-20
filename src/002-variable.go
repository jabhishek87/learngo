package main

import "fmt"

func main() {
	// https://godoc.org/fmt
	// varibale declaration
	a := 10
	b := "golang"
	c := 3.14
	d := true
	fmt.Printf("var a is type %T and val = %v \n", a, a)
	fmt.Printf("var b is type %T and val = %v \n", b, b)
	fmt.Printf("var c is type %T and val = %v \n", c, c)
	fmt.Printf("var d is type %T and val = %v \n", d, d)

	// using shorthand

	var e int = 10
	var f string = "golang"
	var g float64 = 4.17
	var h bool

	fmt.Printf("var e is type %T and val = %v \n", e, e)
	fmt.Printf("var f is type %T and val = %v \n", f, f)
	fmt.Printf("var g is type %T and val = %v \n", g, g)
	fmt.Printf("var h is type %T and val = %v \n", h, h)

	// fmt.Println("hello world")
}
