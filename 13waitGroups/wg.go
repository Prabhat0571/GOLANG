package main

import (
	"fmt"
	"sync"
)

 func task (id int , w *sync.WaitGroup){
	defer w.Done() // ek routine poora hone k bad nikal deta h
	fmt.Println("doing task: ", id)
 }

func main(){
	var wg sync.WaitGroup
	for i:=0 ; i<10; i++{
		wg.Add(1)
		go task(i, &wg)
	
	}
	wg.Wait()
    
}