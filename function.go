package main
import "fmt"
func main(){
	c,d:=MaxMin(50,80)
	fmt.Printf("Max:%d,Min:%d",c,d)
}
func MaxMin(a int,b int)(max int,min int){
	if a>b{
		return a ,b
	}else{
		return b,a
	}
}	