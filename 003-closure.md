# 003-closure


[Prev](002-variable.md) || [Next](004-blank-identifier.md)

```go
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
<
```
To run the program, put the code in hello-world.go and use `go run src/003-closure.go`.