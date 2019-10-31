package main

import "fmt"

var num = make(chan int)
var done = make(chan bool)

func gen(){
	n := 0
	fmt.Println("enter n value")
	fmt.Scan(&n)
	num <- n
}

func cube(){
	n:=<-num
	nn := n*n*n
	fmt.Println(nn)
	done <- true
	
}

func main(){
	go gen()
	go cube()
	//fmt.Println("cube is:", res)
	<-done
}