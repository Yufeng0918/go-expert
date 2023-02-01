package main

import (
	"fmt"
	"go-expert/dao"
	"go-expert/service"
	"go-expert/user"
)

func main() {
	s := user.Hello()
	fmt.Println("Hello, World, " + s)
	service.TestUserService()
	service.Login()
	dao.Insert()
}
