package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, error := os.OpenFile("C:\\Users\\samsung\\Desktop\\test.txt", os.O_CREATE|os.O_RDWR, 0666)
	    if error != nil{
			fmt.Println("error")
		}
	bufrd := bufio.NewReader(file)
	str, err := bufrd.ReadString('/')
	fmt.Println(err)
	fmt.Println(str)
}