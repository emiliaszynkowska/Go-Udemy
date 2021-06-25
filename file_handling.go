package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var name string

	fmt.Print("Name: ")
	fmt.Scan(&name)

	file, err1 := os.OpenFile("names.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0777)
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	file.WriteString("Hello " + name + "\n")

	text, err3 := ioutil.ReadFile("names.txt")
	if err3 != nil {
		fmt.Println("Error:", err3)
	}

	file.Close()

	fmt.Println(string(text))
}
