package main

import (
	"fmt"
)

// slice就是dynamic array, 可以改變長度，用來保存個數不確定的需求
// 他是用來操作陣列的抽象介面，底層是一個struct
// reference type
// iterable,
// var a []int

/* 如何在記憶體保存？
在記憶體會分成三個部分保存，array address, slice本身長度, slice容量大小
所以也能認為slice底層是一個struct

slice := intArr[1:3]   [(intArr[1的address]), len, cap]
intArr [1, 22, 33, 44, 55]

type slice struct {
	ptr *[2]int
	len int
	cap int
}

也就是說，我修改slice data, 也會跟著影響array data, 因為slice儲存的是指向array的指標，直接去修改array memory存放的資料
當我改變slice的長度時，他會分派一個新的記憶體空間，指標指向他，把array的值copy過去，接著把舊的array gc掉
*/

//slice剛宣告完不能使用，一定要引用到一個array才可以使用（因為這樣指標才會指向一塊記憶體位置）

func main() {
	var intArr [5]int = [...]int{1, 22, 33, 44, 55}

	// slice reference 到intArr這個array的第二到第三個元素
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice element: ", slice)
	fmt.Println("slice 個數", len(slice))
	fmt.Println("slice 容量", cap(slice)) //slice的容量可以動態變更 array沒有capibility的概念

	fmt.Println("intArr[1] addr=", &intArr[1])
	fmt.Println("slice[0] addr=", &slice[0])

	makeSlice()
	fmt.Println("--------")
	directAssign()
	fmt.Println("--------")
	iterate()
	fmt.Println("--------")
	sliceAppend()
	fmt.Println("--------")
	sliceCopy()
	fmt.Println("--------")
	practice()
}

func makeSlice() {
	var slice []int = make([]int, 6, 10) // ptr直接指向某個記憶體空間，只能通過slice去操作
	//如果沒有賦值則會給預設值，int, float => 0, string=> "", bool => false
	slice[1] = 10
	slice[2] = 20
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
}

func directAssign() {
	var slice []int = []int{1, 2, 3} // 原理類似make, 只有slice能access到這段記憶體空間
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func iterate() {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice := arr[1:4]

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	for i, v := range slice {
		fmt.Printf("slice[%v]=%v\n", i, v)
	}
}

// slice可以動態增長
func sliceAppend() {
	var slice []int = []int{100, 200, 300}
	fmt.Println("slice", slice)
	slice_new := append(slice, 400, 500, 600) //會分配到另一塊新的記憶體，所以如果要原地追加，就把自己重新賦值就好

	//透過append將slice追加給slice
	slice = append(slice, slice...) //創建新的記憶體，slice將指針重新指向新的記憶體位址
	fmt.Println("slice_new", slice_new)
	fmt.Println("slice: ", slice)

}

func sliceCopy() {
	var slice []int = []int{100, 200, 300}
	var slice_copy []int = make([]int, 10) // 10個element都被初始化為0
	copy(slice_copy, slice)                // copy過後，原本的slice len不變！ 如果copy過來的長度大於本人，他也不會報錯，頂多就是後面的東西不會覆蓋過來
	fmt.Println("slice=", slice)
	fmt.Println("slice_copy", slice_copy)
}

// practice fibonacci
func fbn(n int) []uint64 {
	var fbnSlice []uint64 = make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}

	return fbnSlice
}

func practice() {
	slice := fbn(10)
	fmt.Println("fbn= ", slice)
}
