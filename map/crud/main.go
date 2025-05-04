package main

import (
	"fmt"
)

func main() {
	test := map[string]string{
		"test1": "aa",
		"test2": "bb",
		"test3": "cc",
	}
	fmt.Println(test)

	//因test3已存在，所以以下程式碼等同於修改
	test["test3"] = "dd"

	fmt.Println(test)

	//刪除
	delete(test, "test1")
	fmt.Println(test)

	//如果刪除的key不存在，就不會刪除也不報錯
	//如果要把map的key一次清除，一種方法是interate，一種方法是make一個新的記憶體空間，舊的讓他GC
	delete(test, "test4")
	fmt.Println(test)

	//查找
	val, ok := test["test2"]
	if ok {
		fmt.Printf("find test2!!, vaule= %v", val)
	}

}
