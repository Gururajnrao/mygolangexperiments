package main

import (
    "fmt"
)

func palindrome(input string) string {
    reversed := ""
    for _, char := range input {
        reversed = string(char) + reversed
    }
    return reversed
}
func main() {
    originalstring:="racecar"
    reversestring:=palindrome(originalstring) // Output: "olleh"
    fmt.Println(reversestring)
    fmt.Printf("Original string is same as reversestring:%t",originalstring==reversestring)
}
