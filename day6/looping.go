package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//to make the infinite loop
	for {
		fmt.Println("do stuff infinitely")
	}
	// to stop this loop use ctrl + c

	/**
	x:=5
	for {
		fmt.Println(do stuff)
		if x>10 {
			break
		}
	}

	range operator can be used to iterate over a
	built in ds or user defined one
	*/
}
