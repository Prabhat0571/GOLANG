package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main(){
  fmt.Println("learn channels")
  mychan := make(chan int)
  wg.Add(1)
 
  go passChan(mychan) //reciever first (in case of unbuffered channel reciever should be first otherwise code will get block)
   mychan<-5 //sender later i (else use buffered mychan:= make (chan int , 1))

  wg.Wait()
  
}

func passChan(mychan chan int){
   wg.Done()
   fmt.Println(<-mychan)

}
   