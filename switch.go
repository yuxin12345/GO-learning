package main
import "fmt"
import "time"
func main(){
	i:=5
	switch{
	case i<5:
		fmt.Println("i<5")
	case i==5:
		fmt.Println("i=5")
	case i>5:
		fmt.Println("i>5")
	}
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("Today is a weekend")
	default:
		fmt.Println("Today is a weekday")
	}
	t :=time.Now()
	switch {
	case t.Hour()<12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}
}