package main
import "fmt"
type person struct{
	age int
	name string
}

func main(){
	var slice1[] int =make([]int,10)
	for i := 0; i < 10; i++ {
		slice1[i]=i*5
	}
	fmt.Println(slice1)
	person1 :=new(person)
	person1.age =15
	person1.name ="tthxq"
	fmt.Println(person1)
	map2 :=make(map[string]string,2)
	map2["key1"]="1454"
	map2["key2"]="dfa5"
	fmt.Println(map2)
}