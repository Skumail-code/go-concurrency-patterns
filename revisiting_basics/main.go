package main
import (
    "fmt"
    "slices")

func main() {
  slice := []int{1,2,3,4,5}
  l := len(slice)-1
  
  for i:=l;i>=0;i-- {
    //   for j:=len(slice)-1;j<=0;j--{
          val := slice[i]
          slice = append(slice, val)
        //   fmt.Println(slice)
          slice = slices.Delete(slice, i, i+1)
        //   fmt.Println(slice)
    //   }

  }
  
  fmt.Println(slice)
}
