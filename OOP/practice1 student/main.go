package main

import (
	"fmt"
)

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) Say() {
	infoStr := fmt.Sprintf("student info: name:[%v], gender:[%v], age:[%v], id:[%v], score:[%v]",
		student.name, student.gender, student.age, student.id, student.score,
	)

	fmt.Println(infoStr)
}

func main() {
	student := Student{
		name:   "sophie",
		gender: "female",
		age:    29,
		id:     1,
		score:  100.0,
	}

	student.Say()
}
