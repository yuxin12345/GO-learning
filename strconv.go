package main
import "fmt"
import "strconv"
func main(){
	a:= "go"
	b:=  1
	c:= 5.20
	d := a + strconv.Itoa(b)
	fmt.Printf("%s\n",d)
	f,error:= strconv.ParseInt("go",64,0)
	if error!=nil{
		fmt.Println("err")
	}
	fmt.Printf("%d\n",f)
	g,error := strconv.ParseFloat("go",64)
	if error!= nil{
		fmt.Println("error")
	}
	fmt.Printf("%f\n",g+c)
	x, err := strconv.ParseFloat("-99.7", 64)
	fmt.Printf("%8T %6v %v\n", x, x, err)
	y, err := strconv.ParseInt("71309", 10, 0)
	fmt.Printf("%8T %6v %v\n", y, y, err)
	z, err := strconv.Atoi("71309")
	fmt.Printf("%8T %6v %v\n", z, z, err)
}