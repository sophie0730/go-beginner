package main

import (
	"errors"
	"fmt"
)

func test() int {
	defer func() {
		err := recover() //golang 內建的catch error function
		if err != nil {
			fmt.Println("err: ", err)
		}
	}()
	num1 := 10
	num2 := 1
	res := num1 / num2
	return res
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	}

	return errors.New("read file error") //自定義錯誤
}

func test02() {
	err := readConf("config.in1")
	if err != nil {
		//拋出錯誤並終止程式
		panic("errro!!!!!!")
	}
	fmt.Println("keep continue...")
}

func main() {
	res := test()

	fmt.Println(res)
	test02()
}
