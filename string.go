package main
import "fmt"
func main(){
	str :="hello"
	c:=[]byte(str)
	c[0]='c'
	s2 :=string(c)
	fmt.Println(s2)
    for _,m :=range str{
		fmt.Println(m)
	}
}