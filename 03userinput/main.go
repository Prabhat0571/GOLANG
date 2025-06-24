package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
    
	  reader:= bufio.NewReader(os.Stdin)
	  fmt.Println("Enter the rating for our pizza: ")
	  //comma ok syntax as we dont have try catch in golang
	  input,_:= reader.ReadString('\n')
	  fmt.Println("thanks for rating, " , input) 

}  