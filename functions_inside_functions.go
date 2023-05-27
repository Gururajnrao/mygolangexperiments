// ./main.go
package main

import (
	"fmt"
)

func a(age int) int {
	//fmt.Println("Age", age)
	return age * 2
}

func n(class int) int {
	return class
}

func family(val int, fn func(val1 int) int) {
	result := fn(val)
	//return result
	fmt.Println("Result:", result)
}

func main() {
	age1 := a(41)
	fmt.Println("Age", age1)
	family(41, a)
	//fmt.Println("Result1", result1)

}
