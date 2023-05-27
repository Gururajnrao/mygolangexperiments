// ./main.go
package main

import (
	"fmt"
)

func arrayexample(array []string) {

	fmt.Println("Value inside function:", array[0])
}

func main() {
	//fmt.Println("Bye, asteroid!")
	arrayparam := []string{"guru", "mana", "tropicale"}
	arrayexample(arrayparam)
	// example := [....]string{"guru", "mana", "tropicale"}
	// for i := 0; i < len(example); i++ {
	// 	if i == 1 {
	// 		example[i] = "tower4"
	// 	}
	// 	fmt.Println(example[i])
	// }
}
