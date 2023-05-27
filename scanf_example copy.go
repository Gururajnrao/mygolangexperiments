// ./main.go
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Bye, asteroid!")
	var a int
	var b string
	count, err := fmt.Scanf("%d %s", &a, &b)
	fmt.Println("Value of count", count)
	fmt.Println("Value of error", err)
	fmt.Println("Value of number", a)
	fmt.Println("Value of string", b)
}
