package main

import (
	"fmt"
)

func main() {
	personSalary := map[int]int{
	 3:0,
	 4:0,
	}
	
	fmt.Println("All items of a map")
	for key, _ := range personSalary {
		fmt.Println( key,":",key*key)
	}

}
