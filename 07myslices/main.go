package main
import(
	"fmt"
)

func main()  {
   //slice declaration is just don't mention size
//    var nums[] int 
//    fmt.Println(nums) //returns nil

    var nums = make([]int,0,2 ) //third argument is initial capacity
	// fmt.Println(nums)           //second argument is number of elements to be initialised with zero
	// fmt.Println(cap(nums))
   nums= append(nums, 1,2,3,4,5)
   fmt.Println(nums)
}