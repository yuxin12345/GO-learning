package main
import "fmt"

func main() {
    function1()
}

func function1() {
    fmt.Printf("First\n")
	defer function2()
	defer function3()
	defer function4()
    fmt.Printf("Second\n")
}

func function2() {
    fmt.Printf("last\n")
}
func function3() {
    fmt.Printf("fourth\n")
}
func function4() {
    fmt.Printf("third\n")
}
