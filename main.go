package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go doSomeThing()
	go doSomeThing()
	go doSomeThing()
	time.Sleep(2 * time.Second)
	fmt.Println(time.Since(start))
}

func doSomeThing() {
	time.Sleep(time.Second)
	fmt.Println("something")
}
