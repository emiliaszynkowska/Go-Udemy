package main

import "fmt"

type person struct {
	firstname string
	surname string
	age int
}

type human interface {
	speak()
}

func isInterface(h human) {
	fmt.Println("I am an interface")
}

func (p person) speak() {
	fmt.Println("I am human")
}

func main() {
	myinterface := person {
		firstname : "szyn",
		surname: "szynkowska",
		age : 21,
	}
	myinterface.speak()
	isInterface(myinterface)
}