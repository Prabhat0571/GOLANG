package main
import(
	"fmt"
)

func main()  {
	var fruitList=[]string{"apple","tomato" ,"peach","pineapple","avocado","guava"}
	fmt.Println(fruitList)
	fmt.Printf("type of this data structure is: %T\n ", fruitList)
	fmt.Println(fruitList)
// 	highscore:= make([]int,4)
// 	highscore[0]=345
// 	highscore[1]=645
// 	highscore[2]=745
// 	highscore[3]=1245

// fmt.Println(highscore)
// highscore=append(highscore, 243,234,3545,656)
// fmt.Println(highscore)

var index int =2
fruitList=append(fruitList[:index],fruitList[index+1:]...) //three dots are used to increase the number of arguments in the functions
fmt.Println(fruitList)

}
