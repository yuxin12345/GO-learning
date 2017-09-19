package main
import "fmt"
func SortEven(n int){
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			switch  {
			case i+j<=n &&i<=j:
				fmt.Printf("%d  ",GS1(i,j,n))
			case i+j>=n+2 && i>=j:
				fmt.Printf("%d  ",GS2(i,j,n))
			case i+j<n+2 && j<=n/2:
				fmt.Printf("%d  ",GS3(i,j,n))
			default:
				fmt.Printf("%d  ",GS4(i,j,n))
				
			}
			
		}
		fmt.Println()	
	}
}
func SortOdd(n int){
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			switch  {
			case j>=i&& i<(n+1)/2&&i+j<=n+1:
				fmt.Printf("%d  ",GS1(i,j,n))
			case i+j<=n && j<(n+1)/2:
				fmt.Printf("%d  ",GS3(i,j,n))
			case i==j&&i==(n+1)/2:
				fmt.Printf("%d  ",GS5(i,j,n))
			case i+j>=n+2 &&j>(n+1)/2&&j>i:
				fmt.Printf("%d  ",GS4(i,j,n))
			default:
				fmt.Printf("%d  ",GS6(i,j,n))
				
			}
			
		}
		fmt.Println()
	}
}
func GS1(a int ,b int,c int)(d int){
	 m:=4 *c *(a-1)-(2*b-2)*(2*b-2)+1
	 if a==b {
		 return m
	 }
	 return GS1(a,b-1,c)+1
}
func GS2(a int, b int,c int)(d int){
	if a+b ==c+2{
		m:=GS3(a,b-1,c)-1
		return m
	}
	return GS2(a,b-1,c)-1
}
func GS3(a int,b int,c int)(d int){
	if a==b+1{
		m:=GS1(a,b+1,c)-1
		return m
	}
	return GS3(a-1,b,c)-1
}
func GS4(a int,b int,c int)(d int){
	if a==b-1{
		m:=GS2(a+1,b,c)-1
		return m
	}
	return GS4(a+1,b,c)-1
}

func GS5(a int,b int,c int)(d int){
	m:=c*c
	return m
}
func GS6(a int,b int,c int)(d int){
	if a==b && b> (c+1)/2{
		m:=GS4(a-1,b,c)+1
		return m
	}
	return GS6(a,b+1,c)+1
}

func main(){
	n:=8
	if n%2==0 {
		SortEven(n)
	}
	if n%2==1{
		SortOdd(n)
	}
}