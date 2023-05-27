// ./main.go
package main

import (
	"fmt"
)

type i int

type l struct {
	s int
	h string
}

func (m *l) methodstruct() {
	m.h = "guru1"
	fmt.Println(*m)
}

func (m *i) methodnostruct() {
	*m = 10
	fmt.Println(*m)
}

func main() {

	// val := i(5)
	// val.methodnostruct()
	// fmt.Println(val)

	val1 := l{s: 10, h: "guru"}
	val1.methodstruct()
	val2 := i(5)
	val2.methodnostruct()
	//fmt.Println(val1)

}
