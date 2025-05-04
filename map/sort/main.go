package main

import (
	"fmt"
	"sort"
)

/*
Golang的map是無序的
starting with Go 1.12, the fmt package prints maps in key-sorted order to ease testing.

*/

func main() {
	//但如果要自己排序的話也能這樣做
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 2
	map1[2] = 68
	map1[4] = 57

	fmt.Println(map1)
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map[%v]=%v", k, map1[k])
	}

}
