package main

import (
	"fmt"
)

type MethodUtils struct {
}

// 給MethodUtils寫一個方法
func (mu *MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 邊寫一個方法打印一個m*n的矩形
func (mu *MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 編寫一個方法計算矩形面積（接收長和寬）將其作為方法返回值
func (mu *MethodUtils) Area(width float64, len float64) float64 {
	return width * len
}

// 編寫方法判斷一個數是奇數或偶數
func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 != 0 {
		fmt.Println("奇數")
	} else {
		fmt.Println("偶數")
	}
}

// 根據行、列、指定字符打印相應的效果
func (mu *MethodUtils) Print3(m int, n int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func main() {
	var mu MethodUtils
	mu.Print()
	fmt.Println()
	mu.Print2(2, 6)

	area := mu.Area(5.0, 6.0)
	fmt.Println(area)
	mu.JudgeNum(5)

	fmt.Println()
	mu.Print3(5, 3, "#")
}
