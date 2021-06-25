package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", add(1,1))
}

func add(x int, y int) int {
	return x + y
}