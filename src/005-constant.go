// https://blog.golang.org/constants
package main

import (
	"fmt"
)

func simple_const() {
	const s1 string = "constant string"
	const s2 = "constant string 2"
	const n1 int = 41
	const n2 = 42
	// n2=45;
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
}

func multiple_const() {
	// multiple init
	const (
		PI              = 3.14
		LANGUAGE        = "GO"
		CON      string = "testing"
	)

	fmt.Println("PI=", PI)
	fmt.Println("LANGUAGE=", LANGUAGE)
	fmt.Println("CON=", CON)
}

func iota_const1() {
	const (
		A = iota // 0
		B = iota // 1
		C = iota // 2
	)

    fmt.Println("A=", A)
    fmt.Println("B=", B)
    fmt.Println("C=", C)
}

func iota_const2() {
    const (
        A = iota // 0
        B
        C
    )
    // RESTART FORM 0
    const (
        D = iota // 0
        E
        F
    )

    fmt.Println("A=", A)
    fmt.Println("B=", B)
    fmt.Println("C=", C)
    fmt.Println("D=", D)
    fmt.Println("E=", E)
    fmt.Println("F=", F)
}
// https://www.rapidtables.com/convert/number/binary-to-decimal.html?x=100000000000000000000
func binary_to(){
    const (
        _ = iota // 0
        KB = 1 << (iota * 10) // 1 << (1 * 10)
        MB = 1 << (iota * 10) // 2 << (2 * 10)
        GB = 1 << (iota * 10) // 3 << (3 * 10)
    )
    fmt.Println("Binary to Decimal Conversion")
    fmt.Printf("binary\t %b = %d \n", KB, KB)
    fmt.Printf("binary\t %b = %d \n", MB, MB)
    fmt.Printf("binary\t %b = %d \n", GB, GB)
}

func main() {
	simple_const()
	multiple_const()
    iota_const1()
    iota_const2()
	binary_to()

}
