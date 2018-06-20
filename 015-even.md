# 015-even


[Prev](014-modulo.md) || [Next](016-fizbuzz.md)

```go
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
<
```
To run the program, put the code in hello-world.go and use `go run src/015-even.go`.