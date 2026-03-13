package login

import "fmt"

type UserInfo struct {
	Name string
	Age  int
}


func User() {
	fmt.Println("code by dipesh")
	fmt.Println("This is the customer module.")
	fmt.Println("This is the customer module.")
	fmt.Println("This is the customer module.")
	fmt.Println("This is the customer module.")
	fmt.Println("This is the customer module.")	
	
	user1:= UserInfo{
		Name: "dipesh",
		Age:  21,
	}
	fmt.Println(user1)
}