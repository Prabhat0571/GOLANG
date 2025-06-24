package main

import (
	"fmt"
	"time"
)
func task(id int){
	fmt.Println("doing task", id)
}
func main(){
	fmt.Println("welcome to the goroutines")
	for i:=0; i<=10; i++ {
		time.Sleep(2* time.Second)
		go task(i) //this is a blocking code as it executes sequentially so we write "go"
	}
}
