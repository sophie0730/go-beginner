package main

import "fmt"

func main() {
	users := make(map[string]map[string]string)

	users["tom"] = make(map[string]string, 2)
	users["tom"]["nickname"] = "to"
	users["tom"]["pwd"] = "123"

	modifyUser(users, "tom")
	fmt.Println(users)

	modifyUser(users, "mary")
	fmt.Println(users)

	modifyUser(users, "mike")
	fmt.Println(users)

	modifyUser(users, "mary")
	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {
	_, exist := users[name]
	if exist {
		userInfo := users[name]
		userInfo["pwd"] = "888888"
		return
	}

	newUser := make(map[string]string, 2)
	newUser["nickname"] = name
	newUser["pwd"] = "newpwd"

	users[name] = newUser
}
