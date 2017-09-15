package main
import "fmt"
type Simple interface{
	get() int
}
type set struct{
	x int
}
func (b set) get() int{
	return b.x
}
func inter(g Simple){
	fmt.Println(g)
	fmt.Println(g.get())
}
func main(){
	i:=set{5}
	inter(i)
}