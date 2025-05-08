package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * 3.14 //是從傳入的Circle裡面取的，和main當中的circle是不同的記憶體空間
}

// 為了提高效率，通常method和struct的指標類型綁定，這也代表裡面針對struct的修改，都會影響到原始的instance
func (c *Circle) area2() float64 {

	fmt.Printf("c是 *circle指向的地址%p", c)
	//因為c是指標，所以標準的訪問方式如下
	return (*c).radius * (*c).radius * 3.14
}

func main() {
	var circle Circle
	circle.radius = 4.0
	area := circle.area() //創建獨立的記憶體空間，分別再把c, radius也分配獨立的記憶體空間

	fmt.Println("circle area: ", area)

	var c Circle
	c.radius = 5.0

	//但編譯器底層做了優化，讓我們少寫點程式碼
	fmt.Printf("c address= %p", &c) //這邊的記憶體位址會等同於傳入area2的c的記憶體位址
	c.area2()                       //等同於(&c).area2()
}
