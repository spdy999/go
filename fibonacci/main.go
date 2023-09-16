package main

func fibonacci(max int) {
	fib := make([]int, max)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < max; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
}

func main() {
}
