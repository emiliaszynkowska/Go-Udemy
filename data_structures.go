package main

import "fmt"

func main() {
	var myarray []int
	myarray = []int{1,2,3,4,5}
	myslice := make([]string, 3)
	mymap := map[string]int {
		"simba" : 2004,
		"nala" : 2004,
	}
	type person struct {
		firstname string
		surname string
		age int
	}
	mystruct := person {
		firstname : "simba",
		surname : "priderocks",
		age : 16,
	}

	fmt.Println("Array:", myarray)
	fmt.Println("Slice:", myslice)
	fmt.Println("Map:", mymap)
	fmt.Println("Struct:", mystruct)
}
