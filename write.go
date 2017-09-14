package main

import (
	"strconv"
	"bufio"
	"os"
	"fmt"
)


func main() {
	map1 := make(map[int]string)
    map1[1] = "Monday"
    map1[2] = "Tuesday"
    map1[3] = "Wednesday"
	map1[4] = "Thursday"
	map1[5] = "Firday"
	map1[6] = "Saturday"
	map1[7] = "Sunday"
	var LS string 
	for key, value := range map1 {
		fmt.Println(key,value)
        LS += strconv.Itoa(key) +":" +value +"/   \n"
	}
	LS = LS+ ".end"
	fmt.Println(LS)
	file, error := os.OpenFile("C:\\Users\\samsung\\Desktop\\test.txt", os.O_CREATE|os.O_RDWR, 0666)
	    if error != nil{
			fmt.Println("error")
		}
	bufwriter := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
	bufwriter.WriteString(LS)
	}
	 bufwriter.Flush()
}