// ./main.go
package main

import (
	"fmt"
)

type i int

func (m *i) methodnostruct() {
	*m = 10
	fmt.Println(*m)
}

func main() {

	val := i(5)
	val.methodnostruct()
	fmt.Println(val)
}
