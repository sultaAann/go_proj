package main

import (
	"fmt"
	"os"
)

func AVG() {
	nums := []int{}

	var count int
	fmt.Println("Count of elements:")
	fmt.Scanf("%d\n", &count)
	
	if count <= 0 {
		fmt.Println("n <= 0")
		os.Exit(1)
	}	

	var n int
	for i:= 0; i < count; i++{
		fmt.Scanf("%d\n", &n)
		nums = append(nums, n)
	}

	var total float64 = 0;

	for _, el := range nums {
		total += float64(el)
	}
	fmt.Println(nums)
	fmt.Println("AVG:", total / float64(len(nums)))
}

func Min(arr []int) {
	var min int = arr[0]
	for _, el := range arr {
		if el <= min {
			min = el
		}
	}
	fmt.Println("Min:", min)
}

func SubArr(arr []int, low int, high int) {
	res := arr[low:high]

	fmt.Println(res)
}

func main() {
	//Min([]int{1, 0, 3, 4, 5, 6, 7})
	// AVG()

	SubArr([]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, 5)
}
