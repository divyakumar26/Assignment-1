package main

import "fmt"

func test(str ...string){
  var newstring []string
  for i:=0;i<len(str);i++ {
    if str[i] != "" {
      newstring = append(newstring, str[i])
    }
  }
  fmt.Println(newstring)
}

func main(){
  test("abc","","xyz","mno","")
}