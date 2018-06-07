# 006-memory-address


[Prev](005-constant.md) || [Next](007-pointers.md)

```go
package main

import (
	"fmt"
)

func main() {
	memory_addr1()
	memory_addr2()

}

func memory_addr1() {
	a := 43
	fmt.Println("Value of a is = ", a)
	fmt.Println("Memory address of a is = ", &a)
}

func memory_addr2() {
	var meters float64
	const metersToYards float64 = 1.09361

	fmt.Print("Enter meters swam:")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meters is ", yards, " yards.")
}
<
```
To run the program, put the code in hello-world.go and use `go run src/006-memory-address.go`.