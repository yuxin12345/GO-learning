package main

import (
	"strconv"
	"bufio"
	"os"
	"fmt"
	"io"
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

	file, error := os.OpenFile("C:\\Users\\samsung\\Desktop\\test.txt", os.O_CREATE|os.O_RDWR, 0666)
	    if error != nil{
			fmt.Println("error")
		}
	bufwriter := bufio.NewWriter(file)
	
	bufwriter.WriteString(LS)
	
	bufwriter.Flush()


	file, error = os.OpenFile("C:\\Users\\samsung\\Desktop\\test.txt", os.O_CREATE|os.O_RDWR, 0666)
	if error != nil{
		 fmt.Println("error")
	}
    bufrd := bufio.NewReader(file)
    str, err := bufrd.ReadString('/')
    fmt.Println(err)
    fmt.Println(str)


 inputFile, Error := os.Open("C:\\Users\\samsung\\Desktop\\test.txt")
 if Error != nil {
	 fmt.Printf("error")
	 return 
 }
 defer inputFile.Close()

 inputReader := bufio.NewReader(inputFile)
 for {
	 inputString, readerError := inputReader.ReadString('\n')
	 if readerError == io.EOF {
		 return
	 }
	 fmt.Printf("The input was: %s", inputString)
 }
}