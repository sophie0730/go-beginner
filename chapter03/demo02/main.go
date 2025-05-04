package main

import "fmt"

// 一次定義多個全局變數
var (
	a1 = 10
	a2 = "a"
	a3 = "hi"
)

func main() {
	// 一次聲明多個變數
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	n4, n5, n6 := 1, 2, 3
	fmt.Println(n4, n5, n6)
	fmt.Println(a1, a2, a3)
}
