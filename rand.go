package main
import (
    "fmt"
    "math/rand"
)

func main() {
    for i := 0; i < 10; i++ {
        a := rand.Int()
        fmt.Printf("%d / ", a)
    }
    for i := 0; i < 5; i++ {
        r := rand.Intn(50)
        fmt.Printf("%d / ", r)
    }

}