package main

import "fmt"

func main() {

	var num int = 9
	fmt.Println("num address =", &num)

	var ptr *int = &num
	fmt.Println("ptr =", ptr)

	*ptr = 5
	fmt.Println("modified num:", *ptr)
}
