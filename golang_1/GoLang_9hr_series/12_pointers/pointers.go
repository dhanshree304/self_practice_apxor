package main

//A pointer stores the memory address of another variable.
// Why use pointers?
// ✅ Modify original values
// ✅ Save memory for large data
//*means Go to the address stored in p and fetch the value

import "fmt"

func changeNum(num int){
	num=1
	fmt.Println("In chnagenum",num)//1

}
//.................................


type User struct{
	Name string
}


func main(){
	num :=5
	changeNum(num)//Memory address 0xc00008c0a8
	fmt.Println("Memory address",&num)//it gives the address of num
	fmt.Println("After change",num)//5
	//......................................

	x:=30
p:=&x
fmt.Println("Value of x",x)//30
fmt.Println("Memory address of x :",p)//0xc0000140a0
fmt.Println("Value using pointer",*p)//30  // dereference pointer


name:=User{Name:"Dhanshree"}
ptr:=name //addressing the struct variable

fmt.Println(ptr)//&{Dhanshree}
fmt.Println(ptr.Name) // Go auto-dereferences /Dhanshree


}          




