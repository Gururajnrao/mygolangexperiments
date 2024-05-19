package main
import "fmt"

func fibonacci (n int) int {
  if n<2 {
    return n
  } else {
    return fibonacci(n-1) + fibonacci(n-2)
  }
}

func main() {
  
  n:=5
  fib_output:=fibonacci(n)
  fmt.Println(fib_output)
  
}