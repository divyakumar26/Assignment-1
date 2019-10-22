package main
import "fmt"

func bmi() int{
	var l,b,h int
	sum := 0

	fmt.Println("enter length, breadth and height")
	fmt.Scan(&l,&b,&h)
	sum = l*b*h
	return sum


}

// func main(){
// 	fmt.Println(bmi())
// }