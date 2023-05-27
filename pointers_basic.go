// ./main.go
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Bye, asteroid!")
	a := 10
	address := &a
	*address = 20
	fmt.Println(address)
	fmt.Println(*address)
	fmt.Println(a)

}
