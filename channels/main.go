package main
import "fmt"

func main(){
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
    
    go func(){
        ch1 <- 2
    }()
    
    go func(){
        ch2 <- 4
    }()
    
    go func(){
        ch3 <- 6
    }()
    
    
    select {
        case res1 := <-ch1:
        fmt.Println("Data recievd from ch1", res1)
        case res2 := <-ch2:
        fmt.Println("Data recievd from ch2", res2)
        case res3 := <-ch3:
        fmt.Println("Data recievd from ch3", res3)
    }
}
