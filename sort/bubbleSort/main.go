package main

import (
	"fmt"
)

// bubble sort
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr= ", (*arr))

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}

}

func main() {
	arr := [5]int{24, 65, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println("排序後arr= ", arr)

}
