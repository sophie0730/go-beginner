package main

import "fmt"

// 創建sturct instance的幾種方法
type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {

	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("r1.leftUp.x address=%p\n r1.leftUp.y address=%p\n r1.rightDown.x address=%p\n r1.rightDown.y address=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有兩個*Point類型，這兩個*Point的本身的地址也是連續的，但他們指向的記憶體位址不一定是連續的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp.x address=%p\n r2.leftUp.y address=%p\n r2.rightDown.x address=%p\n, r2.rightDown.y address=%p\n",
		&r2.leftUp.x, &r2.leftUp.y, &r2.rightDown.x, &r2.rightDown.y)

	//但他們指向的地址不一定連續
	fmt.Printf("r2.leftUp point address=%p\n r2.rightDown address=%p\n",
		r2.leftUp, r2.rightDown)
}
