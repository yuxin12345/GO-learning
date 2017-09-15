package main
import "fmt"
type Geometry interface{
	area() int
	perim() int
}
type square struct{
	width int
	height int
}
func (s square) area()int{
	return s.height*s.width
}
func (t square) perim()int{
	return t.height*2+t.width*2
}
func measure(g Geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
func main(){
	s:=square{5,6}
	measure(s)
}