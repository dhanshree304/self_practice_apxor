package main

import "fmt"
import "strings"


func main() {


	var name string = "golang"
	var name1  = "golang1"
	name2:="golang2"
	name1="go"//update variable


	
	
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)

	var isAdult bool = true
	fmt.Println(isAdult)

	var isAdult1  = true
	fmt.Println(isAdult1)

	isAdult2:=false
	fmt.Println(isAdult2)

	
	var price float32 = 50.5
	var price1=50.51
	price2:=50.52

	fmt.Println(price)
	fmt.Println(price1)
	fmt.Println(price2)

const learn string ="react"
const learn1="react1"

// fmt.Println(learn)
// fmt.Println(learn1)
//learn ="golf"//cant update with const 


//multipal constants we canoup together
	const (port =300
		host ="iam")

		// fmt.Println(port)
		// fmt.Println(host)

//----------------------------------------------------------------
//-------------------------------------------------------

s:="hello"
fmt.Println(len(s))
//-------------------------------------------

//✅ 3. Access Characters
fmt.Println(s[0])//o/p 104 ASCII value

//Go strings are byte slices:
fmt.Println(string(s[0])) //will give value at 0th index
//------------------------------------------------------------

//✅ 4. Slicing Strings

str:="Hello,World"
fmt.Println(str[0:3])//Hel
fmt.Println(str[6:])//World

//✅ 6. Compare Strings
// fmt.Println("go"=="GO")//false
// fmt.Println("go"=="go")//true


//✅ 7. Contains / Prefix / Suffix
fmt.Println(strings.Contains("Hello Golang","Gol"))
fmt.Println(strings.HasPrefix("Hello","Hel"))
fmt.Println(strings.HasSuffix("world","ld"))


//✅ 8. Find Index
fmt.Println(strings.Index("Hello","e"))
fmt.Println(strings.LastIndex("Hello","e"))

//✅ 9. Replace
 strr :="Go is cool lang"
fmt.Println(strings.Replace(strr,"cool","awesome",1))
fmt.Println(strings.ReplaceAll(strr,"o","p"))

//✅ 10. Split String ..splits returns to array ✖️
str3:="my Name is D"  
arr:=strings.Split(str3," ")
fmt.Println(arr)

//✅ 11. Join String ✖️
arrt := []string{"Go", "is", "fun"} 
fmt.Println(strings.Join(arrt, " "))

//✅ 12. Uppercase / Lowercase 
fmt.Println(strings.ToUpper("go"))//GO
fmt.Println(strings.ToLower("GO"))//go

//✅ 13. Trim Spaces
fmt.Println(strings.TrimSpace("   hello  ")) //"hello"
fmt.Println(strings.Trim("....g.g.....","."))//g.g
}