package main

import "fmt"

// 匿名函數就是只使用一次就丟掉的函數

// 也可以定義全局的匿名函數
var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20) //n1 = 10, n2 = 20

	fmt.Println(res1)

	res4 := fun1(10, 20)
	fmt.Println(res4)
}
