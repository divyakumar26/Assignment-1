package main

import "fmt"

func palindrome() string{
	var str string

	fmt.Println("enter a string")
	fmt.Scan(&str)
	for i:=0;i<len(str)/2;i++ {
      if str[i] == str[len(str)-i-1] {
		  return "true"
	  }
	}
	      return "false"

}

// func main(){
// 	fmt.Println(palindrome())
// }