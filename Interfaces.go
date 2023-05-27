// ./main.go
package main

import (
	"fmt"
)

type declaration interface {
	cost() int
}

type car struct {
	brand string
	price int
}

type bike struct {
	brand string
	price int
}

func (c car) cost() int {
	c.price = c.price + 10000
	return c.price
}

func (c bike) cost() int {
	c.price = c.price + 1000
	return c.price
}

func printcost(d declaration) {
	fmt.Println(d.cost())

}

func main() {

	// val := i(5)
	// val.methodnostruct()
	// fmt.Println(val)

	c := car{brand: "Dzire", price: 500000}
	b := bike{brand: "Honda", price: 50000}
	printcost(c)
	printcost(b)

}
