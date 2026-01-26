package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
 
func getStatusCode(endpoint string){

	defer wg.Done()
	res,err := http.Get(endpoint)

	if err!=nil {
		fmt.Println("response not coming \n", err )
	} else{
		fmt.Printf("%d status code for %s \n",res.StatusCode,endpoint)
	}

}



func main(){

	

	websiteList:= []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://go.dev",

	} 	

	for _, web :=range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

