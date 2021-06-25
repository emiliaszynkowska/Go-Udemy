package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var count int

func main() {
	waitGroup.Add(3)
	go hello()
	go world()
	go increment("count")
	waitGroup.Wait()
}

func hello() {
	fmt.Println("hello")
	waitGroup.Done()
}

func world() {
	fmt.Println("world")
	waitGroup.Done()
}

func increment(s string) {
	for i:=0; i<10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		count++
		fmt.Println(s, i, "Counter:", count)
		mutex.Unlock()
	}
	waitGroup.Done()
}