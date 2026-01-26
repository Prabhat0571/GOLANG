package main

import (
	"fmt"
	"sync"
)
 
var wg sync.WaitGroup

type post struct{
	mu sync.Mutex
	views int
}

func (p* post)inc(){
	wg.Done()
	p.mu.Lock()
	p.views+=1
	defer p.mu.Unlock()
}

func main(){
	myPost:= post{views:0}
	for i:=0; i<100; i++{
		wg.Add(1)
	   go myPost.inc()
	}
	wg.Wait()
	fmt.Println(myPost.views)


}