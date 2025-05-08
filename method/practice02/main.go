package main

import (
	"fmt"
)

// 實現一個計算機功能，有加減乘除四個功能
type Calculator struct {
	Num1 float64
	Num2 float64
}

func (cal *Calculator) Calculate(operator string) float64 {
	res := 0.0
	switch operator {
	case "+":
		res = cal.Num1 + cal.Num2
	case "-":
		res = cal.Num1 - cal.Num2
	case "*":
		res = cal.Num1 * cal.Num2
	case "/":
		res = cal.Num1 / cal.Num2
	}

	return res
}
func main() {
	calculator := Calculator{
		Num1: 10,
		Num2: 20,
	}
	res := calculator.Calculate("-")
	fmt.Println(res)
}
