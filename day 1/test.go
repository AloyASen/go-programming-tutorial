package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
this is a multi line comment

this is just a note about how function nameing is in go

the functions that have their initial letter capitalized
is assured to be exported by the package in the compiler
allowing other tools to get to it

the functions that do not start with caps are made internal
to the package

 ---- @ this was as described by sentdex in his tutorials

*/
func foo() {
	fmt.Println("the  square root of 4 is ", math.Sqrt(4))
	fmt.Println("a number from 1 to 100", rand.Intn(100))
}
func add(x, y float64) float64 {
	return x + y
}

const a int = 23

func multiple(a, b string) (string, string) {
	return a, b
}
func main() {
	//foo()
	var num1, num2 float64 = 5.6, 9.5
	fmt.Println(add(num1, num2))
	// the numbers can be declared aroud dynamically but once fixed
	//they cannot be changing their types

	//if there is unused variables in go they will throw errors
	// not warnings also any unused import is reported to behave
	//the same [therfore both of the following lines are commented together]

	//num3, num4 := 4.5, 3.2
	//fmt.Println(add(num3, num4))

	//multiple return from one function
	fmt.Println(multiple("hey there ", "this is aloy"))

	//exlicit typecasting
	var a int = 52
	var b float64 = float64(a)
	fmt.Println("the converted value of the function ", b)

	//type inference also works
	// x :=a --> this will make x with the type int
	//this code is commented as it is not used anywhere in code
}
