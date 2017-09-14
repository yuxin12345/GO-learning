package main
import "fmt"
func main() {
	i:=1
	for i<=3{
		fmt.Println("i:",i)
		i++
	}
	for j := 4; j < 7; j++ {
		fmt.Println("j:",j)
	}
	for {
		fmt.Println("next")
		break
	}
	k:=8
	if k<10 {
		fmt.Println("true")
	}
	if k<5{
		fmt.Println("k<5")
	}else{
		fmt.Println("k>=5")
	}
	if k<5{
		fmt.Println("k<5")
	}else if k>10{
		fmt.Println("k>10")
	}else{
		fmt.Println("5<=k<=10")
	}
}