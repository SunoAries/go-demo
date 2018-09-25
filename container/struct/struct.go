package main

import "fmt"

type User struct {
	Name string
	Password string
	NickName string
	Age  int32
	mess string
}

func main() {

	var user User
	var user1 *User = &User{}
	var user2 *User = new(User)
	fmt.Println(user, *user1, *user2)
}
