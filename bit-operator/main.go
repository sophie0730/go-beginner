package main

import "fmt"

func main() {

	var i int = 5
	fmt.Printf("%b \n", i) //二進制

	var j int = 011 // 011 => 9
	fmt.Println("j= ", j)

	var k int = 0x11 // 17
	fmt.Println("k= ", k)

	operator()
	bitOperator()
}

func operator() {
	// 2 & 3
	// 0000 0010
	// 0000 0011
	// 0000 0010 => 2
	fmt.Println(2 & 3)

	// 2 | 3
	// 0000 0010
	// 0000 0011
	// 0000 0011 => 3
	fmt.Println(2 | 3)

	// 2 ^ 3 相當於XOR, 如果兩個其中一個為1就為1 否則就是0
	// 0000 0010
	// 0000 0011
	// 0000 0001 => 1
	fmt.Println(2 ^ 3)

	// -2 ^ 2
	// 原數 1000 0010 但要使用補數 由1111 1101 + 1 =
	// 所以-2的二進制為1111 1110
	// 0000 0010
	// 1111 1100（補數） => 原數為 1111 1100 - 1 = 1111 1011 > 1000 0100 => -4
	fmt.Println(-2 ^ 2)
}

func bitOperator() {
	// >> 為右移，<<為左移
	// a := 1 >> 2, 0000 0001 => 0000 0000 = 0(就直接擠掉)
	// c := 1 << 2, 0000 0001 => 0000 0100 = 4
	a := 1 >> 2
	c := 1 << 2
	fmt.Println(a)
	fmt.Println(c)
}
