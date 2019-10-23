package main

import "fmt"

func main(){
	words := []string{"abc","xyz","abc","lnm", "abc"}
	for i:=0;i<len(words);i++ {
		for j:=i+1;j<len(words);j++ {
			if words[i] == words[j] {
               words = append(words[:i],words[i+1:]...)
			}
		}
	}

	
	fmt.Println(words)
}