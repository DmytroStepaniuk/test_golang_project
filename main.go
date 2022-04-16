package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1
	fmt.Println("Hi!   ")

	fmt.Println(math.Exp2(float64(a)))

	fmt.Println("Bye!   ")

	SayHello()
}

func SayHello() {
	fmt.Println("Hello!")
}
