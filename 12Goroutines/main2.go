package main
import (
	"fmt"
)



func main(){
	for i:=0; i<6; i++ {
		greeter("hello")
		greeter("world")
	}
}

func greeter(s string){
	fmt.Println(s)

}