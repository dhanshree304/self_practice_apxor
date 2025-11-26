package main

import (
	"fmt"
	
)

func main() {

	//on array
	intTings := []int{1, 2, 3, 4, 5}
	fmt.Println(intTings)
	
	strThings:=[]string{"aai","tai","mai"}
	for i:=0; i<len(strThings); i++ {
	fmt.Println(strThings[i])
}

// In Go, range gives two values:
// Index (position)
// Value (actual item)


for _,i:=range strThings{//range word will iteratre overate each elem in strThings
	fmt.Println(i)
}


//while loop in go
start:=1
for start <=100{
	start+=start
	if start ==32{
		break
	}
	fmt.Println("Start is now :",start)
}


//iteratre on string

str:="golang"
for i,v:=range str{
	fmt.Println(i,v)
}

//iterate on object
student :=map[string]int{
	"no":1,
	"price":30,
	"lane":2,
}



for key,value:=range student{
	fmt.Println(key,":",value)
}
}