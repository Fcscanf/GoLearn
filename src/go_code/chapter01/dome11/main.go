package main

import "fmt"

func main() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"

	modifyUser(users, "Tom")
	modifyUser(users, "mary")

	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "88888"
		users[name]["nickname"] = "昵称~" + name
	}
}
