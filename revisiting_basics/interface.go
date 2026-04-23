package main
import "fmt"

type Bird interface{
    Fly() string
}

type Pegion struct{
    Color string
}

func (p *Pegion) Fly() string {
    return fmt.Sprintf("%s pegion is flying", p.Color)
}

func main() {
    p := Pegion{
        Color: "Grey",
    }
    
    fmt.Println(p.Fly())
}
