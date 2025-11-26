package main

import "fmt"

func main() {
// arr := [5]int{1, 2, 3, 4, 5}
// // slice1 := arr[1:4] //give index here 
// // fmt.Println(slice)//[2 3 4]

// slice1 := make([]int, 3) // len=3, cap=3
// fmt.Println(slice1)

s1 := []int{1, 2, 3}
s2 := make([]int, len(s1))
copy(s2, s1)
fmt.Println(s2)

}
