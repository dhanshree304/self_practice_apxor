package main

import "fmt"

func main() {
    var x, y int

    fmt.Print("Enter two numbers: ")
    fmt.Scan(&x, &y)

    fmt.Printf("x = %d, y = %d\n", x, y)
    fmt.Printf("Sum = %d\n", x+y)
}
