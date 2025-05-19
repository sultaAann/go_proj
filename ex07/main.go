package main

import (
	"fmt"
	"os"
	"time"
)

func wr(c chan<- int) {
	for i := 0; ; i++ {
		c <- i
	}
}

func rd(c <-chan int) {
	for {
		num := <-c 
		if num % 2 == 0 {
			fmt.Println(num)
		}
	}
}

func timer(n int) {
	select {
	case <- time.After(time.Duration(n) * time.Second):
		fmt.Println("Timer")
	}
}

func main() {
	//cn := make(chan int)

	//go wr(cn)

	//go rd(cn)
	//n := 0
	//fmt.Scanf("%d", &n)
	

	timer(3)

	os.Exit(0)
}
