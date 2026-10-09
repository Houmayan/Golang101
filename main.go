package main

import (
	"fmt"
	"go-program/test"
)

func main() {
	fmt.Println("Hello Go!")

	test.Test()
	test.VariableExample()
	test.ArrayPractice()
	test.SlicePractice()
	fmt.Println(test.PracticeFunction(1,3,3,4,5,6,6))
	fmt.Println(test.PracticeFunction(1,2,2,2,2,4))
	result := test.PracticeClosure()
	fmt.Println("closure")
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())

	result2 := test.PracticeClosure()
	fmt.Println("second Closure")
	fmt.Println(result2())
	fmt.Println(result2())
	fmt.Println(result2())
	fmt.Println(result2())
	fmt.Println(result2())
	fmt.Println("First closure Value Check:")
	fmt.Println(result())
}