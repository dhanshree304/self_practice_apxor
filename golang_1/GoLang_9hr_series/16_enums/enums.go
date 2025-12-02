package main

import "fmt"

//enumerated types
//Go uses const + custom types + iota to create enum-like behavior.


func changeOrderStatus(status string){
	fmt.Println("changing order status to",status)
}

func main(){
changeOrderStatus("Confirmed")


}