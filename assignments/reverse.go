package main
import "fmt"

func reverse(){
	var n string 
	fmt.Println("enter n value")
	fmt.Scan(&n)
	for i:=len(n)-1;i>=0;i-- {
		fmt.Print(string(n[i]))
		


	}
}

// func main(){
// 	reverse()
// }