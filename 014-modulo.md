# 014-modulo


[Prev](013-hello-name.md) || [Next](015-even.md)

```go
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
<
```
To run the program, put the code in hello-world.go and use `go run src/014-modulo.go`.