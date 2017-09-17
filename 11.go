package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e10)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- (i*4 + 1)
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- (i*4 + 3)
	}
}

func suck(ch1, ch2 chan int) {
	var sun int
	for v := range ch1 {
		sun += v
		fmt.Println( sun)
	}
	for v := range ch2 {
		sun += v
		fmt.Println( sun)
	}
}