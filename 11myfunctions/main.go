package main
import(
	"fmt"
)
// func add(a int , b int) int{
// 	return a+b
// }

// //to return a function with multiple values
// func languages()(string, string, string){
//       return "python", "rust" , "javascript"
 // }

func processit(add func(a int)int){
	add(2)

}
func main(){
	fmt.Println("hello from go")
// 	 result2, result3 , result4 := languages()
// 	 fmt.Println( result2, result3, result4)
      
    }



