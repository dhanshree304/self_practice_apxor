package main

import "fmt"

func main() {
	//array initializes itself with 0 if we dont give it values
	//zeroed values

	var nums [4]int
	fmt.Println(len(nums),nums)//4 [0 0 0 0]

	var vals [4]bool
	vals[1]=true
	fmt.Println(vals)

	var str [4]string
	str[1]="a"
	fmt.Println(str)



	//2d array
	nums2:=[2][2]int{{3,4},{5,6}}
	fmt.Println(nums2)//[[3 4] [5 6]]
}
