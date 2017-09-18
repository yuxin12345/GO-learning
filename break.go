package main
import "fmt"
func main(){
	k:=15
	for {
		k--
		if k<0{
			break//for无线循环，需要break结束
		}
		fmt.Printf("%d",k)
		fmt.Printf(" ")
	}
	fmt.Println()
	for i:=0; i<3; i++ {
        for j:=0; j<10; j++ {
            if j>5 {
                break   
            }
            print(j)
        }
        print("  ")
    }
}