package main

import "fmt"

func main() {
	var mystring string
	var myint int8
	var myfloat float32
	var mybool bool

	mystring = "hello world"
	myint = 1
	myfloat = 0.1
	mybool = true

	fmt.Println("String:", mystring)
	fmt.Println("Int:", myint)
	fmt.Println("Float:", myfloat)
	fmt.Println("Bool:", mybool)
}
