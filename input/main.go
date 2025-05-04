package main

import "fmt"

func main() {

	var name string
	var age byte
	var sal float32
	var isPass bool

	fmt.Println("Please enter name")
	fmt.Scanln(&name) //相當於是把輸入的字串傳入到name所在的address存入

	fmt.Println("Please enter age")
	fmt.Scanln(&age)

	fmt.Println("Please enter salary")
	fmt.Scanln(&sal)

	fmt.Println("Please enter if passing test")
	fmt.Scanln(&isPass)

	fmt.Printf("%s %d %f %t", name, age, sal, isPass)

}
