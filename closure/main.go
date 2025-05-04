package main

import (
	"fmt"
	"strings"
)

// 可以把閉包想成是class, function式method, n是字串。函數和他使用的n構成一個閉包
// 反復調用函數時，n只初始化一次，因此每調用一次會進行累加
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x // n 被內部函數引用，記憶體被分配到heap段，所以function執行完畢，n的記憶體不會被回收
		return n
	}
}

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, ".jpg") {
			return name + ".jpg"
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(3)) //14
	fmt.Println(f(5)) //19

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("winter"))
	fmt.Println(f2("bird.jpg"))
}
