# 013-hello-name


[Prev](011-conditions.md) || [Next](014-modulo.md)

```go
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
<
```
To run the program, put the code in hello-world.go and use `go run src/013-hello-name.go`.