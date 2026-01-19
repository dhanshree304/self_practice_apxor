package main

import "fmt"

//enumerated types
//Go uses const + custom types + iota to create enum-like behavior.

type OrderStatus string //for single type 
//for multiple types we use struct

type UserRole string
type User struct {
	Name string `json:"name"`
	Role UserRole `json:"role"`
}

const 
(
Received OrderStatus ="received your order"
Confirmed OrderStatus  = "confirmed your order"
Prepared OrderStatus = "prepared your order"
Delivered OrderStatus  = "delivered your order"

)

const (
	Admin UserRole="admin"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to",status)
}

func main(){
changeOrderStatus(Received)

// create user with enum role
u:=User{
	Name:"Dhanshree",
	Role:Admin,
}
fmt.Println(u)
}

