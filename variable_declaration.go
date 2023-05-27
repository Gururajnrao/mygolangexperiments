// ./main.go
package main

import (
  "fmt"
  "./data"
)

func main() {
  //fmt.Println("Bye, asteroid!")
  fmt.Println(data.Message)
}

/* Initialize Package data */

package data

Message := "This is demo message"

