// ./main.go
package main

import "fmt"

func main() {
	//fmt.Println("Bye, asteroid!")
	example := [...]string{"guru", "mana", "tropicale", "Sarjapur"}
	slice := example[1:4]
	fmt.Println("Slice:", slice)
	fmt.Println("Length of Slice:", len(slice))
	fmt.Println("Cap of Slice:", cap(slice))

	// Subslice
	slice_1 := slice[1:]
	fmt.Println("Subslice:", slice_1)
	fmt.Println("Length of SubSlice:", len(slice_1))
	fmt.Println("Cap of SubSlice:", cap(slice_1))

	//Append to a slice
	slice_3 := append(slice_1, "Bangalore", "Carmelaram", "Dinne")
	fmt.Println("Slice append:", slice_3)
	fmt.Println("Length of Slice append:", len(slice_3))
	fmt.Println("Cap of slice append:", cap(slice_3))

	//Delete from Slice
	slice_4 := slice_3[3:]
	fmt.Println("Slice4:", slice_4)
	delslice := append(slice, slice_4...)
	fmt.Println("Slice delete:", delslice)
	fmt.Println("Length of Slice delete:", len(delslice))
	fmt.Println("Cap of slice delete:", cap(delslice))

	// Copying from a slice
	dest_slice := make([]string, 2)
	val := copy(dest_slice, slice_4)
	fmt.Println("Dest slice:", dest_slice)
	fmt.Println("Values in dest slice:", val)
}
