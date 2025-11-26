package main

import (
	
	"fmt"
	"slices"
)

//in go , a slice is a flexible and dynamic data structure that provides a more powerful alternative to array.
//slices have dynamic size and their length can be changed during the programs execution.

func main(){
numbers :=[]int{1,2,3,4,5}
numbers=append(numbers,3, 6,7,8,9)
fmt.Println("Numbers:",numbers)//Numbers: [1 2 3 4 5 3 6 7 8 9]
fmt.Printf("Numbers has data type: %T\n",numbers)//printf used to print data type
fmt.Println("Length :",len(numbers))//Length : 10


//create slice using make() fun
//if we add elem more than capacity it will add them in slice and increase (doubles) the capacity
numbers2 := make([]int,3,4) //taking len,capacity
numbers2=append(numbers2,4)
numbers2=append(numbers2,5)
numbers2=append(numbers2,6)
fmt.Println(numbers2)

fmt.Println("slice:",numbers2)
fmt.Println("Length:",len(numbers2))
fmt.Println("Capacity:",cap(numbers2))
fmt.Println(numbers2==nil)//false

var nums1=[]int{1,2,3}
fmt.Println(nums1[:2])//[1]
fmt.Println(nums1[0:1])//[1 2]
fmt.Println(nums1[1:])//[2 3]

var nums2=[]int{4,5,6}
slices.Reverse(nums2)
fmt.Println(nums2)//[6 5 4]

clonedNums:=slices.Clone(nums2)
fmt.Println(clonedNums)


//2d slices:=
d2slice:=[][]int{{1,2,3},{4,5,6}}
fmt.Println(d2slice)//[[1 2 3] [4 5 6]]

}