package main

import "fmt"

func main() {
	x := 3

	for i:=0; i<10; i++ {
		fmt.Print(i)
	}

	if x > 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	switch x {
	case 1:
		fmt.Println(0)
	case 2:
		fmt.Println(1)
	case 3:
		fmt.Println(2)
	}
}
