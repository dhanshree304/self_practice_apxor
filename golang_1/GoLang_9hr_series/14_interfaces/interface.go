package main

import "fmt"

//interfaces means collection of methods signatures
//it makes relation with multiple methods

//An interface defines a set of methods. Any type that implements those methods satisfies the interface.

//..........................................
type values struct {
first int
second int

}

type mathTest interface{
add(a,b int) int
multiply(a,b int)int
//divide(a,b int)int
//substraction(a,b int)int
}
func (v values) add(a, b int) int{
return a+b + v.first +v.second
}


func (v values) multiply(a, b int) int{
return a*b *v.first * v.second
}
//................................................

//1)Define interface
type Shape interface{
	Area() float64
}
//2)Create struct
type Rectangle struct{
	height float64
	width float64
}
type Circle struct {
	radius float64
}

//functions takes inputs of struct and method of interface
//3) Implement methods
func (r Rectangle) Area() float64{
	return r.width * r.height
}

func (c Circle) Area() float64{
return 3.14 *c.radius *c.radius
}


func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
func main (){
//create struct object 

reaVal:=values{
	first:2,
	second:3,
}
//assign interface to real values of struct
var m mathTest =reaVal

fmt.Println("Add Result:",m.add(5,5))
fmt.Println("Multiply Result:",m.multiply(2,3))
//............................................


realVRect := Rectangle{10, 5}
realValCir := Circle{7}

printArea(realVRect)
printArea(realValCir)


}