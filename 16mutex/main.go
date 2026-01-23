package main

import (
	"fmt"
	"sync"
)

type post struct {
	mu sync.Mutex
	views int
}  

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	//mutex 
	p.mu.Lock() //bnd kia
	p.views += 1 //modify hua
	p.mu.Unlock() //khol dia

}

func main() {
    var wg sync.WaitGroup 
	myPost := post{views: 0}
	for i:= 0 ; i<100 ; i++{
        wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()

	fmt.Println(myPost.views)
	
}