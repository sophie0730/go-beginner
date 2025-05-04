package main

import (
	"fmt"
)

//引入別的package時：
//被引用的變量定義先執行
//在執行import package的init
//最後才是main文件的init

// 全局變數會第一個執行
var age = test()

func test() int {
	fmt.Println("test")
	return 90
}

// init函數在每個源文件當中可以有一個，可以在main執行之前完成初始化的工作

func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main")
	fmt.Println("age", age)
	fmt.Println(utils.age, utils.name)
}
