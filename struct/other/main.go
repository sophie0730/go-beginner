package main

import "fmt"

// 要如何把一個instance的type轉成另一個struct?
// 兩個sturct裡面的變數要完全一樣，才可以進行轉換
type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {

	var a A
	var b B

	a = A(b)
	fmt.Println(a, b)
}
