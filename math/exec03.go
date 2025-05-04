package main

import "fmt"

func main() {
	var a int = 10
	var b int = 5
	a = a + b
	b = a - b // b = a + b - b
	a = a - b // a = a + b - a - b + b

	fmt.Println(a, b)
}
