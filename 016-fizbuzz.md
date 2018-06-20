# 016-fizbuzz


[Prev](015-even.md) || [Next](README.md)

```go
package main

import (
	"fmt"
)


/*
Write a program that prints the numbers from 1 to 100. 
But for multiples of three print "Fizz" instead of the number and 
for the multiples of five print "Buzz". For numbers which are 
multiples of both three and five print "FizzBuzz".
*/

func main() {
	fmt.Print("Please Enter a UpperLimit: ")
	var upperLimit int
	fmt.Scanf("%d", &upperLimit)

	for i:=0; i<=upperLimit; i++ {
		// fmt.Println(i)
		if i%15 == 0{
			fmt.Println(i, "FizBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
<
```
To run the program, put the code in hello-world.go and use `go run src/016-fizbuzz.go`.