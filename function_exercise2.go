package main
import "fmt"
import "math"
func main(){
	fmt.Println(MySqrt(0.64))

}
func MySqrt(a float64)(float64,err error){
	if a<0{
		return fmt.Println("err")
	}else{
	return math.Sqrt(a),nil
	}
}