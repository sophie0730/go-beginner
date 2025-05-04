package main

import (
	"fmt"
	"strings"
)

func test(b byte) byte {
	return b + 1
}

func main() {
	//go 中的switch區塊一次只會執行一個
	//所以不需要另外加break
	// case後面只要是返回一個值都可以，所以可以是function也可以是表達式

	// var key byte
	// fmt.Println("please enter a char, a,b,c,d,e,f,g")
	// fmt.Scanf("%c", &key)

	// switch key {
	// case 'a':
	// 	fmt.Println("Monday")
	// case 'b':
	// 	fmt.Println("Tuesday")
	// case 'c':
	// 	fmt.Println("Wendnesday")
	// default:
	// 	fmt.Println("Not suitable day")
	// }

	// switch 'a' {
	// case 'a':
	// 	fmt.Println("Monday")
	// case 'b':
	// 	fmt.Println("Tuesday")
	// case 'c':
	// 	fmt.Println("Wendnesday")
	// default:
	// 	fmt.Println("Not suitable day")
	// }

	// switch test(key) {
	// case 'a':
	// 	fmt.Println("Monday")
	// case 'b':
	// 	fmt.Println("Tuesday")
	// case 'c':
	// 	fmt.Println("Wendnesday")
	// default:
	// 	fmt.Println("Not suitable day")
	// }

	var n1 int32 = 20
	var n2 int32 = 32

	switch n1 {
	case n2, 10, 5: //error, n1 和n2要是同個data type 並且case後面可以以逗號表達多個條件
		fmt.Println("ok")
	}

	// 也能當作if else使用
	var age int = 10

	switch {
	case age == 10:
		fmt.Println("10")
	case age == 20:
		fmt.Println("20")
	}

	//根據範圍做判斷
	var score int = 30
	switch {
	case score > 90:
		fmt.Println("great")
	case score >= 80 && score < 90:
		fmt.Println("20")
	}

	//switch 的穿透，最多只能穿透一層，
	var num int = 10
	switch {
	case num < 50:
		fmt.Println("ok")
		fallthrough
	case num == 10:
		fmt.Println("okok")
		fallthrough //沒辦法穿透到default
	case num < 40:
		fmt.Println("falied")
	default:
		fmt.Println("default")
	}

	practice()
}

func practice() {
	var character byte
	fmt.Println("please enter a char, a,b,c,d,e")
	fmt.Scanf("%c", &character)

	switch character {
	case 'a', 'b', 'c', 'd', 'e':
		fmt.Println(strings.ToUpper(string(character)))
	default:
		fmt.Println("Other")
	}

}
