package main
import "fmt"
func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            continue
        }
        fmt.Printf("%d",i)
        fmt.Printf(" ")
    }
}