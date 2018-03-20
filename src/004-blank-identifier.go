package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func bi01() {
	a := "stored in a "
	// b := "stored in b "

	fmt.Println("a- ", a)

	// fmt.Println("b- ", b)
	// b is not being used
}

func without_bi() {
	// here in this code we are using err which might be tha case that
	//we dont use it
	res, err := http.Get("http://www.amazon.in/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

func with_bi() {
	// here we are using blank indentifier for err
	// underscore is used for blank identifier
	res, _ := http.Get("http://www.amazon.in/")

	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", page)
}

func main() {
	// bi01()
	// without_bi()
	with_bi()
}
