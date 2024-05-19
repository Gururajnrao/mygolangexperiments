// Consider we have num1=2, num2=3. To swap these two numbers, we can just write: num1,num2 = num2, num1
// [1 2 3] change it as [3 2 1]

package main
import "fmt"

func swap(arr []int) []int {
  
  newarr := make([]int, len(arr))
  for i:=0;i<len(arr);i++ {
    newarr[len(arr)-1-i]=arr[i]
  }
  
  return newarr
  
}

func main() {
a:=[]int{1,2,3}
swaparr:=swap(a)
fmt.Println("Values of swapped array: ",swaparr)  
}