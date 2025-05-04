package main

import "fmt"

func main() {
	practice1()
	practice2()
}

func practice1() {
	var total int = 0
	var count int = 0

	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			total += i
			count += 1
		}
	}
	fmt.Println(total)
	fmt.Println(count)
}

func practice2() {
	for i := 0; i <= 6; i++ {
		fmt.Printf("%d + %d = %d\n", i, 6-i, 6)
	}
}
