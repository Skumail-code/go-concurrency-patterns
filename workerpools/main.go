package main
  
import (
    "fmt"
    )

func worker(id int, jobs <-chan int, res chan <- int){
        for job := range jobs {
                res <- job * 2
        } 
}

func main(){
        jobs := make(chan int, 5)
        res := make(chan int, 5)

        for i:=1;i<=3;i++{
                go worker(i, jobs, res)
        }
        
        for j:=1;j<=5;j++{
            jobs <- j
        }
        
        close(jobs)
        
        for r:=1;r<=5;r++{
            fmt.Println(<-res)
        }



}



