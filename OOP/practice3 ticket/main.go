package main

import "fmt"

//一個景點根據遊客年齡收取不同價格門票，年齡大於18收費100元，其餘門票免費
//寫出Visitor struct, 根據年齡決定能夠購買的門票價格並輸出

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) buyTicket() string {
	if visitor.Age > 18 {
		return "100"
	}
	return "free"
}

func main() {
	visitor := Visitor{
		Name: "sophie",
		Age:  29,
	}

	res := visitor.buyTicket()

	fmt.Printf("%v should buy ticket using $%v", visitor.Name, res)
}
