package main

import (
	"fmt"
)

/*
* go當中的map是無序的
* key不能夠重複
* 必須使用make分配記憶體空間，才可以使用
* reference type
 */

func main() {
	var a map[string]string         //宣告不代表分配記憶體，必須使用make
	a = make(map[string]string, 10) //分配10個放key-valuek的記億體空間
	a["no1"] = "hi"
	a["no2"] = "hello"
	fmt.Println(a)

	//也能這樣賦值
	test := map[string]string{
		"test1": "aa",
		"test2": "bb",
	}
	fmt.Println(test)

	students := make(map[string]map[string]string)

	students["stu01"] = make(map[string]string, 2)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "male"

	students["stu02"] = make(map[string]string, 2)
	students["stu02"]["name"] = "Mary"
	students["stu02"]["sex"] = "female"

	fmt.Println(students)

}
