package main
import "fmt"
func main(){
	fmt.Println(add(5,6))
	c:=sub(12,5)
	fmt.Println(c)
	d:=mul(2,4)
	fmt.Println(d)
}
func add(a int,b int)int{
	return a+b
}
func sub(a int,b int)int{
	return a-b
}
func mul(a int,b int)int{
	return a*b
}
