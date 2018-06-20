# 010-dictionaries


[Prev](009-lists.md) || [Next](011-conditions.md)

```go
package main

import (
	"fmt"
)

// https://stackoverflow.com/questions/41668053/cap-vs-len-of-slice-in-golang

func main() {
	go_map := make(map[string]int)
	go_map["first_key"] = 25

	go_map["2nd_key"] = 30

	// remove key
	delete(go_map, "first_key")

	if element, ok := go_map["first_key"]; ok {
		fmt.Println(element) // won't be printed
	}

	if element, ok := go_map["2nd_key"]; ok {
		fmt.Println(element)
	}

}<
```
To run the program, put the code in hello-world.go and use `go run src/010-dictionaries.go`.