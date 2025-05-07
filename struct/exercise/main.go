package main

import "fmt"

// 創建sturct instance的幾種方法
type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person //一個記憶體空間，包含name, age
	p1.Age = 10   //分配值
	p1.Name = "tome"

	var p2 *Person = &p1    //p2是一個指標，指向儲存p1 struct的記憶體空間
	fmt.Println((*p2).Name) //直接去修改p1 struct記憶體空間
	(*p2).Age = 20

	fmt.Printf("p1 age: %v, p1 name: %v, p2 age: %v, p2 name: %v", p1.Age, p1.Name, p2.Age, p2.Name)

}
