package main

import (
	"fmt"
)

func main() {
	x := 15
	a := &x //this is the memory address
	fmt.Println(a)
	fmt.Println(*a)
	*a = 5
	fmt.Println(*a)
	//multiply using this as pointer indirects
	fmt.Println(*a * *a)
}
