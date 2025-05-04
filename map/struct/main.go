package main

import "fmt"

/*
* map的value時常是struct type
 */

type stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	students := make(map[string]stu, 10)

	stu1 := stu{"tom", 10, "Taiwan"}
	stu2 := stu{"Mary", 10, "Taiwan"}

	students["stu01"] = stu1
	students["stu02"] = stu2

	fmt.Println(students)
}
