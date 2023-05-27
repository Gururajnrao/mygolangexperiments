// ./main.go
package main

import (
  "fmt"
  "./data"
)

func main() {
  fmt.Println(data.Message)
}

// ./data/data.go
package data

Message := "This is message from data package"
