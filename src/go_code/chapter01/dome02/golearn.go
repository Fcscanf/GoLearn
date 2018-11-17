package main

import "fmt"

func modifyUser(user map[string]map[string]string, name string) {
	if user[name] != nil {
		user[name]["password"] = "123"
	} else {
		user[name] = make(map[string]string, 2)
		user[name]["password"] = "123"
		user[name]["nickname"] = "昵称" + name
	}
	return
}

func main() {
	users := make(map[string]map[string]string, 10)
	modifyUser(users, "tom")
	modifyUser(users, "fc")
	fmt.Println(users)
	fmt.Println("Hello,Go")
}
