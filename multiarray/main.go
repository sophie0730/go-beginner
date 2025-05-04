package main

import (
	"fmt"
)

/*二維陣列的記憶體分配
arr2   ptr | ptr
arr2[0] address:0x1400012c030
arr2[1] address:0x1400012c048
相差24 byte，因為每個元素有3個int,一個int是8 byte

會有一個記憶體儲存ptr, 這個ptr會指向另一塊記憶體，存放值
如果我是arr[2][3], 那就會分派兩塊記憶體，每塊記憶體儲存三個element

*/

func main() {
	var arr [3][4]int
	arr[1][2] = 1
	arr[2][3] = 3
	fmt.Println(arr)

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Printf("arr2[0] address:%p\n", &arr2[0])
	fmt.Printf("arr2[0] address:%p\n", &arr2[1])

	fmt.Printf("arr2[0][0] address:%p\n", &arr2[0][0])
	fmt.Printf("arr2[0][1] address:%p\n", &arr2[0][1])
	traverse()
}

func traverse() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
	}
	fmt.Println()

	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v", i, j, v2)
		}
	}
}
