package main
//here we use 2 diff func for two diff inputes
//with one hoe we can solve it is defined by generics
import "fmt"


func printIntSlice(items []int){
	for _,item:= range items {
		fmt.Print(item,",")
	}
	fmt.Println()//1,2,3,4,
}

func printStringSlice(items []string){
	for _,items:=range items{
		fmt.Println(items)
	}
}
//a
//b
//c
func printStringSlice1(items []string) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println() // newline at the end
}//a b c


func printAnyType[T any](items []T){
	for _,items:= range items{
		fmt.Println(items)
	}
}
//a
//b
//c
//or
//1
//2
//3
func printMultipleType[T int|bool|string](items []T){
	for _,items:= range items{
		fmt.Println(items)
	}
}

func main(){

	nums:=[]int {1,2,3,4}
	names:=[]string {"a","b","c"}
	bVal:=[]bool {true,false,false,true}
	printIntSlice(nums)
		//printIntSlice(names)//it will not work for names as this func takes slice of
		//int as input

printStringSlice(names)


printStringSlice1(names)//a b c

printAnyType(nums)
printAnyType(names)



printMultipleType(bVal)
//true
// false
// false
// true


printMultipleType(nums)
// 1
// 2
// 3
// 4


printMultipleType(names)
// a
// b
// c

}