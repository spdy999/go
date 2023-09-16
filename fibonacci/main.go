package main

import "fmt"

func fibonacci(max int, ch chan int) {
	fib := make([]int, max)
	fib[0] = 0
	fib[1] = 1

	ch <- fib[0]
	ch <- fib[1]

	for i := 2; i < max; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		ch <- fib[i] // send(?)
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibonacci(6, ch)

	for msg := range ch {
		fmt.Println(msg)
		// result
		// 0
		// 1
		// 1
		// 2
		// 3
		// 5
	}
}
