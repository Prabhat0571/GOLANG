package main
import(
	"fmt"
)
func main() {
	fmt.Println("hello welcome to maps")
	list:= make(map[int]string)
	list[1]= "dal"
	list[2]= "chawal"
	list[3]= "bhindi"
	list[4]= "arbi"
	fmt.Println(list)
     delete(list, 1)
	 fmt.Println(list)
	
	

	
}