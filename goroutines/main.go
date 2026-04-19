package main
import(
	"fmt"
	"time"
)

func run(){
	for i := 0; i <= 5; i++{
		fmt.Printf("The current loop is %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go run()
	time.Sleep(1 * time.Second)
}

