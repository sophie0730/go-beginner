package main

import "fmt"

func main() {
	fibo := fibonacci(4)
	fmt.Println(fibo)

	fmt.Println(getFuncion(1))
	fmt.Println(getFuncion(2))
	fmt.Println(getFuncion(3))
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

/*
已知f(1)=3; f(n) = 2*f(n-1)+1; 求f(n)
*/
func getFuncion(n int) int {
	if n == 1 {
		return 3
	}

	return 2*getFuncion(n-1) + 1
}
