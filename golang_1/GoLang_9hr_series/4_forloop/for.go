package main
import "fmt"

func main(){

	for i:=0; i<=3;i++{
		fmt.Println(i)
		if i==2{
			continue;
		}
	}

//...............................


	// for i:= range 11{
	// 	fmt.Println(i)
	// }


//...............................



//var age int=10✔️
// var age =10✔️ 

	age:=13

	if age >=18{
		fmt.Println("Person is adult")
	}else if age >=12 && age <=17{
		fmt.Println("Person is teenager")
	}else{
		fmt.Println("Person is kid")
	}


	if age :=20;age>=18{
		fmt.Println("Person is adult",age)
	}else if age >=12 && age<=17{
		fmt.Println("Person is teenager",age)
	}else{
		fmt.Println("Person is Kid")
	}
}