package main

import (
	"fmt"
)

func main(){
	mySlice:= []string{"Apple" , "banana" , "pomegrenate"}
	for oneSlice, values:= range mySlice{
		fmt.Println(oneSlice,values)
	}
}

