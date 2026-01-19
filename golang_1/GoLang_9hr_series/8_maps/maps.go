package main

import "fmt"

//maps-->objects

func main(){

	m:=make(map[string]int)//means obj has string type key and int value
		m["price"]=30000
	m["age"]=20

	delete(m,"price")
	fmt.Println(m)//map[age:20]

	clear(m) //removing all key val pairs in our obj/map
fmt.Println(m)//map[]



//..............................................

ma:=map[string]int{"price":110,"no":200}
fmt.Println(ma)//map[no:200 price:110]

vals,ok:=ma["price"]//(1st is return value,2nd return boolean)
fmt.Println(vals)//110
if ok {
	fmt.Println("all ok")//printed as price key is there
}else{
	fmt.Println("not ok")
}




}