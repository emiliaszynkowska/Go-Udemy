package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("none.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
