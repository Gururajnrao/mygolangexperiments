package main
import (
  "fmt"
  "sync"
)

func producer(arr [] string, ch chan string, wg *sync.WaitGroup) {
  for _,value:=range arr {
    defer wg.Done()
    fmt.Println("Producing:",value)
    ch<-value
    //time.Sleep(500*time.Millisecond)
  }
}

func consumer (ch chan string, wg *sync.WaitGroup) {
  for value:=range ch {
    defer wg.Done()
    fmt.Println("Consuming:",value)
    //time.Sleep(500*time.Millisecond)
  }
}



func main() {
  array_str:=[]string{"Tower4","T3","mana","tropicale"}
  var wg sync.WaitGroup
  ch:=make(chan string)
  go producer(array_str, ch, &wg)
  wg.Add(4)
  go consumer(ch, &wg)
  //wg.Add(4)

  wg.Wait()
  
  //time.Sleep(5000*time.Millisecond)
  fmt.Println("Main function execution completed")
}
