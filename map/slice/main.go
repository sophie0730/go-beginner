package main

import "fmt"

/*
slice的類型可以是map，這樣map個數就能動態變化了
*/

func main() {
	//宣告一個slice
	animals := make([]map[string]string, 2) //分配記憶體位置, 2個map

	if animals[0] == nil {
		animals[0] = make(map[string]string, 2)
		animals[0]["name"] = "cat"
		animals[0]["age"] = "3"
	}

	//使用append函數動態增加slice, 就不受一開始定義的記憶體空間所影響
	newAnimal := map[string]string{
		"name": "dog",
		"age":  "4",
	}

	animals = append(animals, newAnimal)

	fmt.Println(animals)
}
