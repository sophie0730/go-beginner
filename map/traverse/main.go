package main

import "fmt"

func main() {
	test := map[string]string{
		"test1": "aa",
		"test2": "bb",
		"test3": "cc",
	}

	for k, v := range test {
		fmt.Printf("k= %v, v=%v\n", k, v)
	}

	fmt.Println("test has ", len(test), "key-value")

	students := make(map[string]map[string]string)

	students["stu01"] = make(map[string]string, 2)
	students["stu01"]["name"] = "tom"
	students["stu01"]["sex"] = "male"

	students["stu02"] = make(map[string]string, 2)
	students["stu02"]["name"] = "Mary"
	students["stu02"]["sex"] = "female"
	fmt.Println("students have ", len(students), "key-value")

	for k1, v1 := range students {
		fmt.Printf("k1= %v", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk2=%v,v2= %v", k2, v2)
		}
	}
}
