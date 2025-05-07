package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 給Person struct添加speak方法
func (person Person) speak() {
	fmt.Println(person.Name, "is a good man")
}

func (person Person) calculate() {
	result := 0
	for i := 1; i <= 1000; i++ {
		result += i
	}

	fmt.Println(person.Name, "calculate result is ", result)
}

func (person Person) calculate2(n int) {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}

	fmt.Println(person.Name, "calculate result is ", result)
}

func (person Person) getSum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	var p Person
	p.Name = "sophie"

	p.speak()
	p.calculate()
	p.calculate2(10)
	res := p.getSum(10, 20)
	println("res= ", res)
}
