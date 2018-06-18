package main

import "fmt"

func main() {
	switch_cond(1.0)
}

function switch_cond(value){
	switch value {
	case 1.0: 
		fmt.Println("value 1.0")
	case 2.0: 
		fmt.Println("value 2.0")
	case 3.0: 
		fmt.Println("value 3.0")
	}
}