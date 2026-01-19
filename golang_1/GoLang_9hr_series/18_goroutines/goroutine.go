package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task",id)
}

func main() {

	for i:=0;i<=10;i++{
	go task(i)
go func (){
		fmt.Println(i)//i is using closure here 
	}()
	}
time.Sleep(time.Second*2)//if dont add that line go routine would not run
//Without this line, the program would finish immediately and goroutines would NOT get time to run.
}