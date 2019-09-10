package main

import "fmt"

func main(){
	fmt.Println("hello world")
	var a chan int = make(chan int)
	a<-5
	v:=<-a
	fmt.Println(v)
}
