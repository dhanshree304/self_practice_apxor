//Value Receiver vs Pointer Receiver() most imp


//methods on structs

package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Birthday() {
    p.Age++
}
func (p *Person) MyBirthday() {
    p.Age++
}

func main() {
    p := Person{Name: "Dhanshree", Age: 22}

	//1) A copy of the struct passed==value receiver
    p.Birthday()
    fmt.Println(p.Age)//22 as p is only the copy 


	//2) Pointer to struct is passed==pointer receiver(original value modified)
	p.MyBirthday()
	fmt.Println(p.Age)//23 as original value is modified
}


