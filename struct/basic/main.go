package main

import "fmt"

// 用struct實踐OOP
type Cat struct {
	Name  string //大寫開頭 代表public
	Age   int
	Color string
}

// 如果struct的類型是: 指標、slice、map，預設值都是nil，代表還沒分配記憶體空間
// 要先使用make, 才能使用
type Person struct {
	Name   string
	Age    int
	Scores [5]*float64       //array
	ptr    *int              //指標
	slice  []int             //slice
	map1   map[string]string //map
}

func main() {
	//創建Cat variable (類似new)
	var cat1 Cat
	cat1.Name = "Una"
	cat1.Age = 3
	cat1.Color = "white"
	fmt.Println(cat1)

	//定義Person struct variable
	var p1 Person
	fmt.Println(p1)

	//使用slice, 要先make一個記憶體空間出來，才可以賦值
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	fmt.Println(p1.slice)

	//使用map, 一定要先make一個記憶體空間出來
	p1.map1 = make(map[string]string)
	p1.map1["name"] = "tom"
	fmt.Println(p1.map1)

	//不同struct變量的字段是獨立互不影響的，一個struct variable的更改不影響另一個，因為他是value type, not reference type
	var cat2 Cat
	cat2.Name = "Mary"
	cat2.Age = 10
	cat2.Color = "black"

	cat3 := cat2 //struct式value type, 為value copy, 除非我直接把cat2 address傳給cat3, 讓cat3指向cat2記憶體位址，這才會變reference type
	cat3.Name = "tom"

	fmt.Println("cat2= ", cat2) //cat2=  {Mary 10 black}
	fmt.Println("cat3= ", cat3) //cat3=  {tom 10 black}

}
