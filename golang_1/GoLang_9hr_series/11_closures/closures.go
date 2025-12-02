
package main
import "fmt"


//fun anonymous is present inside counter fun
func counter() func()int{
	var count int =0
return  func()int{
	count+=1
	return count
}

}


func main(){
increment :=counter()
fmt.Println(increment())//1
fmt.Println(increment())//2

}



