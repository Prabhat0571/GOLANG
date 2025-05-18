package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
     fmt.Println("Welcome to our pizza app")
	 reader:=bufio.NewReader(os.Stdin)
	 fmt.Println("Hi Please Rate Us: ")
	 input,_:= reader.ReadString('\n')
	 fmt.Println("Thanks for rating us: ", input)
	 numRating ,err := strconv.ParseFloat(strings.TrimSpace(input) ,64)
	 if (err!=nil){
		fmt.Println(err)
	 } else{
		fmt.Println("added 1 to your rating: " , numRating+1)
	 }
}