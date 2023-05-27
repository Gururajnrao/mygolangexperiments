// ./main.go
package main

import (
	"fmt"
)

type s struct {
	name     string
	pincode  int
	pincode1 int
}

func structfn(val *s) {
	fmt.Println(*val)
	pincode1 := val.pincode + 1
	val.pincode1 = pincode1

}

func main() {

	c := s{name: "guru", pincode: 560035}
	structfn(&c)
	fmt.Println(c)

}
