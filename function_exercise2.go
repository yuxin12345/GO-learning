package main
import "fmt"
import "math"
func main(){
	fmt.Println(MySqrt(0.64))
}
func MySqrt(a float64)(float64){
	return math.Sqrt(a)
}