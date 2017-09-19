package main

import "fmt"

func main() {
    result := 0
    for i := 0; i <= 11; i++ {
        result = xx(i)
        fmt.Printf("xx(%d) is: %d\n", i, result)
    }
}

func xx(n int) (res int) {
if n==0{
	return 1
}else{
	return n*xx(n-1)
}
}