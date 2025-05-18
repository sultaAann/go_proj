package main

import (
	"fmt"
	"os"
)

func Sum(arr []int) int {
	var total int = 0
	for _, el := range arr {
		total += el
	}
	return total
}

func Half(num int) (int, bool) {
	if num%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

func Max(args ...int) int {
	var res int
	for _, el := range args {
		if res < el {
			res = el
		}
	}
	return res 
}

func Fib(n int) int {
	if (n == 0) {
		return 1
	}

	return n * Fib(n - 1)
}

func MakeOdd() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	//fmt.Println(Sum([]int{1, 2, 3, 4, 5, 5, 6}))
	//fmt.Println(Half(5))
	//fmt.Println(Max(1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 8, 65, 46, 646, 35, 535, 35, 74, 23, 325, 999))
	//fmt.Println(Fib(3))
	
	//next := MakeOdd()
	//other := MakeOdd()
	//fmt.Println(next())
	//fmt.Println(other())

	//fmt.Println(other())
	//fmt.Println(other())

	//fmt.Println(next())
	
	


	os.Exit(0)
}
