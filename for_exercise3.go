package main
import "fmt"
func main(){
	for i := 0; i < 100; i++ {
		switch  {
		case i%3==0 && i%5 !=0:
			fmt.Println("Fizz")//3的倍数且不是5的倍数
		case i%5==0 && i%3 !=0:
			fmt.Println("Buzz")//5的倍数且不是3的倍数
		case i%5==0 &&i%3==0:
			fmt.Println("FizzBuzz")//3的倍数且是5的倍数
		default:
			fmt.Println(i)
		}
	}
}