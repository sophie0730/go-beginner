package main

import "fmt"

// 創建sturct instance的幾種方法
type Person struct {
	Name string
	Age  int
}

func main() {
	//method 1

	//method2
	p2 := Person{}
	fmt.Println(p2)

	//method3
	var p3 *Person = new(Person) //創建struct, 讓p3指向struct的記憶體位址
	//因爲p3是一個指標，因此標準的賦值方式為
	(*p3).Name = "smith"
	(*p3).Age = 30
	//但也可以p3.Name, 只是為了使用方便，但他其實是在底層是做了優化
	//會給p3加上取值運算，實際上p3他還是一個指標
	p3.Age = 32
	fmt.Println(*p3)

	//method4
	var person *Person = &Person{}
	//因為person是一個指標，因此標準的訪問方法
	// (*person).Name = "scott"
	//但為了使用方便也可以直接person.Name = "scott"
	//原因和上面一樣
	(*person).Name = "scott"
	(*person).Age = 88
	fmt.Println(*person) //去找person指向的記憶體位址的值
}
