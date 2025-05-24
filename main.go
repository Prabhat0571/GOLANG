package main
import (
	"fmt"
)
func main() {
	fmt.Println("Welcome to structs")
	//this are the alternative of the classes as we don't have classes in the golang
	//we also don't have INHERITANCE in the golang
	type User struct {
		Name   string
		email  string
		status bool
		age    int
	}
	
	hitesh:= User{"prabhat", "prabhat@go.dev",true, 22 }
	fmt.Println(hitesh)
}
