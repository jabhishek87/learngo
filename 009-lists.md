# 009-lists


[Prev](008-loop.md) || [Next](010-dictionaries.md)

```go
package main

import (
	"fmt"
	"strings"
)

// https://stackoverflow.com/questions/41668053/cap-vs-len-of-slice-in-golang

func main() {
	var arr [3]int // [0,0,0]
	arr[0] = 125  // chnage the element via index

	// arr = append(arr, 340) cannot do append as it length is fixed 3
	colours := []string{"res", "blue"} // slices when size is unknown

	colours = append(colours, "green") // append == recreation

	fmt.Println("length arr = ", len(arr)) // length
	fmt.Println("capacity arr = ", cap(arr)) // capacity
	fmt.Println("value arr=", arr)
	
	fmt.Println(strings.Repeat("--", 20))
	
	fmt.Println("length arr = ", len(colours)) // length
	fmt.Println("capacity arr = ", cap(colours)) // capacity
	fmt.Println("value arr = ", colours)

}<
```
To run the program, put the code in hello-world.go and use `go run src/009-lists.go`.