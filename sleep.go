package main
import "fmt"
import "time"
func main(){
	t :=time.Now()
	fmt.Println(t)
	time.Sleep(time.Second*5)
	fmt.Println(t)
}