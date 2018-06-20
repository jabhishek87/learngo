# 008-loop


[Prev](007-pointers.md) || [Next](009-lists.md)

```go
package main

import "fmt"

func main() {
	for i:=0; i<=100; i+=2 {
		fmt.Println(i)
	}

	// Range loop

	names := []string{
		"Abhishek",
		"Priti",
		"Dhrishti",
	}

	for i, name  := range names {
		fmt.Printf("%d, %s\n", i, name)
	}
}
<
```
To run the program, put the code in hello-world.go and use `go run src/008-loop.go`.