package main

import (
	"fmt"
	"os"
)

func FizzBuzz() {
	fmt.Println("Write number n > 0")
	var n int

	fmt.Scanf("%d\n", &n)
	if n <= 0 {
		fmt.Println("n <= 0; n should be n > 0")
		os.Exit(1)
	}
	for i := 1; i <= n; i++ {
		if i % 15 == 0 {
			fmt.Print("FizzBuzz")
			continue
		}
		if i % 3 == 0 {
			fmt.Print("Fizz ")
			continue
		} else if i % 5 == 0 {
			fmt.Print("Buzz ")
			continue
		} else {
			fmt.Printf("%d ", i)		
		}
		
	}
	
}

func FarToCel() {
	fmt.Println("Write degrees in Far")

	var n float64

	fmt.Scanf("%f\n", &n)
	
	n = (n - 32) * 5/9

	fmt.Println("Result in Cel:", n)
}

func FootToMeter() {
	fmt.Println("Write foot")

	var n int

	fmt.Scanf("%d\n", &n) 

	fmt.Println("Foot in meters:", float64(n) * 0.3048)
}

func DivToThree() {
	fmt.Println("Write number n > 0") 

	var n int
	fmt.Scanf("%d\n", &n)
	
	if n <= 0 {
		fmt.Println("n <= 0; n should be n > 0")
		os.Exit(1)
	}
	for i := 1; i <= n; i++ {
		if i % 3 == 0 {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	DivToThree()
}
