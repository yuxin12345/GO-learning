package main
import "fmt"
func main() {
	nums := []int{2,3,4}
	sum:=0
	for _, num :=range nums{
		sum += num
	}
	fmt.Println("sum:",sum)//sum=2+3+4
	for i,_:=range nums{	
		fmt.Println("index:",i)//index:0,1,2,3,4,5...
	}
	for _,num :=range nums{
		fmt.Println("num:",num)
	}
	map1 := map [string]string{"a":"apple","b":"banana"}
	for key , value := range map1{
		fmt.Printf("%s:%s\n", key , value )
	}
	for i,c :=range "hello"{
		fmt.Println(i,c)
	}
}