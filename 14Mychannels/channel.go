package main

import (
	"fmt"
	"time"
)

func processNum(numchan chan int){  //other go routine 
	fmt.Println("number is: ", <-numchan)
    
}

func main(){
   numchan:= make(chan int)
   go processNum(numchan) //another go routine
   numchan <- 5
   time.Sleep(time.Second*2) 
       

}
   