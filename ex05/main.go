package main

import (
	"fmt"
	"os"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b *int
	c := 5
	d := 6
	a = &c
	b = &d

	fmt.Println(*a, *b)

	swap(a, b)

	fmt.Println(*a, *b)

	os.Exit(0)
}
