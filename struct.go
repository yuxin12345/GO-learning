package main
import "fmt"
type person struct{
	name string
	age  int
	address string
	sex string
}
func main(){
	yu := person{"yuxin",22,"China","man"}
	fmt.Println(yu)
	Yx := person{name:"yuxin",age:22,address:"China",sex:"man"}
	fmt.Println(Yx)
	var Yxin person
	Yxin.name ="yuxin"
	Yxin.age=22
	Yxin.address="China"
	Yxin.sex="man"
	fmt.Println(Yxin)
}