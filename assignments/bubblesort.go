package main

import (
"fmt"
)

func bubble() []int{
 
  a := []int{5,2,9,1}
  for i:=0;i<len(a);i++ {
  for j:=0;j<len(a)-i-1;j++ {
  if a[j] > a[j+1] {
temp := a[j]
a[j]=a[j+1]
a[j+1]=temp
}
}
}
return a
}

func main() {
fmt.Println(bubble())
}
