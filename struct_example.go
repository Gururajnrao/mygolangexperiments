// ./main.go
package main

import (
	"fmt"
)

type district struct {
	city       string
	party      string
	population int
}

func main() {
	//fmt.Println("Bye, asteroid!")
	s := district{
		city:       "Coimbatore",
		party:      "DMK",
		population: 20000,
	}
	fmt.Println(s)

}
