package main

import "fmt"

//maps-->objects

func main(){

	m:=make(map[string]int)//means obj has string type key and int value
		m["price"]=30000
	m["age"]=20

	delete(m,"price")
	fmt.Println(m)

	clear(m)
fmt.Println(m)



//.......................

ma:=map[string]int{"price":110,"no":200}
fmt.Println(ma)

vals,ok:=ma["price"]//(1st is return value,2nd return boolean)
fmt.Println(vals)//110
if ok {
	fmt.Println("all ok")//printed price key is there
}else{
	fmt.Println("not ok")
}




}