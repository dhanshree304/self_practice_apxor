package main

import "time"
import "fmt"

//A struct is a user-defined data type that groups different variables (fields) under one name.

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

//struct with slice and map
type Company struct {
    Name    string
    Employees []string
    Offices   map[string]string //keys:string value:string
}
//nested structs 
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Address Address
}

//🔹 Struct with Functions (Methods)

type Rectangle struct {
    Length float64
    Width  float64
}

func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}



func main() {
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	myOrder2:=order{
		id:"2",
		amount:100,
		status:"delivered",
		createdAt:time.Now(),
	}

	company:=Company{
		Name:"Dhanshree",
		Employees:[]string{"er","jh"},
		Offices:map[string]string{"as":"ty","ty":"Gh"},
	}
	fmt.Println(company)

e := Employee{
    Name: "Ravi",
    Address: Address{
        City:  "Mumbai",
        State: "Maharashtra",
    },
}


rect := Rectangle{
	Length: 10, 
	Width: 5,
}
fmt.Println("Area:", rect.Area())

var rect2 Rectangle
rect2.Length = 10
rect2.Width = 5
fmt.Println("area:",rect2.Area())
fmt.Println(rect)
fmt.Println(e.Address.City)


	myOrder.createdAt = time.Now()
	myOrder2.createdAt=time.Now()
	fmt.Println(myOrder.status)
	fmt.Println("myOrder2 struct",myOrder2)//myOrder2 struct {2 100 delivered {13997240475037836664 1 0x2e4980}}
	fmt.Println("Order struct",myOrder)//Order struct {1 50 received {13997240475037836664 1 0x2e4980}}

}