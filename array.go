// ./main.go
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Bye, asteroid!")
	example := [...]string{"guru", "mana", "tropicale", "chikkanayakana"}
	for i := 0; i < len(example); i++ {
		if i == 1 {
			example[i] = "tower4"
		}
		fmt.Println(example[i])
	}
}
