package main

import (
	"fmt"
	"time"
)
//structs helps us to achieve the functionality of classes as we dont have them in go
type order struct{
	id string
	amount int
	rating int 
	status string
	createdAt time.Time //timestamp

}

func (o *order) changeAmount(amount int){ //pass by reference k lye * lgao 
	o.amount=amount
}

func main(){
	
	nayaorder:= order{
		id: "1",
		amount: 200,
		rating: 5,
		status: "pending",
		
	}

	// doosraOrder:= order{
	// 	id: "1",
	// 	amount: 20320,
	// 	rating: 325,
	// 	status: "confirmed",
	// }

	nayaorder.createdAt= time.Now()
	
	// fmt.Println(nayaorder)
	// fmt.Println(doosraOrder)


	nayaorder.changeAmount(300);
	fmt.Println(nayaorder)

}

