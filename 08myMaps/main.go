package main
import(
	"fmt"
)
func main() {
	fmt.Println("hello welcome to maps")
	languages:=make(map[int]string)
	languages[1]="english"
	languages[2]="hindi"
	fmt.Println(languages)
	fmt.Println(languages[1])
	//to remove we use delete
	delete(languages,1)
	fmt.Println(languages)

	
}