//function in go have ability to return multiple values
//so we need to store this values in different variables

package main

import (
	"fmt"
	
)


func getLan()(string,string,string){
	return "js","py","c";
}
func superman(){
	fmt.Println("I am super girl")
}
func multiply(v1 int,v2 int)int {
	return v1*v2

}
//A variadic function is a function that can accept zero or more arguments of the same type.
//variadic functions
func adder(values ...int) int{
	sum:=0

	for i:=range values{
		sum=sum+values[i]
	}
	return  sum
}

func addmore(values...int)(int,int,string){
sum :=0
for i:=range values{
	sum=sum+values[i]

}
length:=len(values)
name:="Just for fun"
return sum,length,name
}


func sums(nums...int)int{
	total:=0

	for _,num:=range nums{
		total =total+num
	}
	return total
}



func printAll(values ...interface{}){
    for _, v := range values {
        fmt.Println(v)
    }
}
	

func printAll1(values ...interface{}) []interface{} {
    return values
}


// func printfir(values...interface{}) interface{}{
// 	if len(values)==0{
// 		return nil
// 	}
// 	return values[0]
// }



func main(){
	lang1, lang2,lang3:=getLan()
	fmt.Println(lang1,lang2,lang3)

	superman()

	result:=multiply(2,3)
	fmt.Println(result)

	fmt.Println(adder(3,4,5,6))

	fmt.Println(addmore(2,3,4,5,6,7))

	//fmt.Println(sums(2,3,4,5,"hello"))//not work give error

	numbers:=[]int{1,3,4,5}
	fmt.Println(sums(numbers...))// add(numbers) ❌ wrong

	printAll(1, "Hello", 3.14, true)//prints all values each on single line

	fmt.Println(printAll1(1, "Hello", 3.14, false))//[1 Hello 3.14 false]

	//fmt.Println(printfir())//<nil>
	//fmt.Println(printfir(2,4,5,6))//2
}
