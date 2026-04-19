package main
import (
    "fmt"
    "sync")
    
func run(wg *sync.WaitGroup) {
    defer wg.Done()
    for i:=0; i<5; i++ {
        fmt.Printf("Current loop running %d \n", i)
    }
}

func main() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    go run(&wg)

    wg.Wait()
}
