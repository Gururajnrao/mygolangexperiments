// ./main.go
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Bye, asteroid!")
	multiexample := [2][3]int{{1, 1, 3}, {1, 2, 3}}
	fmt.Println(len(multiexample))
	for i := 0; i < len(multiexample); i++ {
		for j := 0; j < len(multiexample); j++ {
			fmt.Println(multiexample[i][j])
		}
	}
}
