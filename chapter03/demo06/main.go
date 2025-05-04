package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//查看某個變數的類型與佔用的記憶體空間
	var n2 int64 = 10

	fmt.Printf("n2 type is %T use %d bytes memory space", n2, unsafe.Sizeof(n2))

	// 程式能正確運行的情況下，會盡量使用較少的記憶體
	var age byte = 10 //byte: 0-255
	fmt.Println(age)

	//golang浮點數會默認是float64
	var num6 = 0.12
	fmt.Printf("type is %T, use %d bytes", num6, unsafe.Sizeof(num6))

	num8 := 5.1234e2   //5.12434 * 10^2
	num9 := 5.1234e2   //5.12434 * 10^2
	num10 := 5.1234e-2 //5.12434 / 10^2
	println(num8, num9, num10)
}
