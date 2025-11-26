package main

import (
	"fmt"
	"time"
)

func main(){

//simple switch
	i:=5

	switch i {

	case 1 :
		fmt.Println("line one")
	case 2 :
		fmt.Println("line two")
	case 3 :
		fmt.Println("line three")
		default :
		fmt.Println("Other")
	}


//multiple condition switch

switch time.Now().Weekday(){
case time.Saturday,time.Sunday:
	fmt.Println("its weekend")
default: fmt.Println("its workday")
}


//type switch

whoAmI:=func(i interface{}){
	switch t := i.(type){
	case int:
		fmt.Println("its a integer")
	case string :
fmt.Println("its a string")
	case bool :
		fmt.Println("its a boolean")
	default:
		fmt.Println("other",t)

	}
}


whoAmI("golang")


}







