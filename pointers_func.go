// ./main.go
package main

import (
	"fmt"
)

func modifyval(val *int) {

	*val = 20
	fmt.Println("Value inside function", *val)
}

func modifymap(loc map[string]string) {
	loc["city"] = "Mysore"
}

func modifyarr(arr []int) {
	arr[0] = 40
}

func main() {
	//fmt.Println("Bye, asteroid!")
	// a := 10
	// modifyval(&a)
	// fmt.Println(a)
	// Map and modify MAP
	// location := make(map[string]string)
	// location["name"] = "Guru"
	// location["city"] = "Bangalore"
	// location["state"] = "Karnataka"
	// modifymap(location)
	// fmt.Println(location)
	// Array, slice and modify
	array := []int{10, 20, 30}
	modifyarr(array)
	fmt.Println(array)
	slice := array[1:3]
	modifyarr(slice)
	fmt.Println(slice)
}
