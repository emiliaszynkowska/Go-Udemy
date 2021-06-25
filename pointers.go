package main

import "fmt"

func main() {
	x := 1
	p := &x

	fmt.Printf("\n%v has pointer %v", x, &x)
	fmt.Printf("\n%v has pointer %v", *p, p)
}
