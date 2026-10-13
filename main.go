package main

import (
	"fmt"
	train "go-program/Train"
	"go-program/test"
)

func main() {

	trainName := []string{"ICE","RE","S-Bahn","Bus","Euro City"}
	trainDelaytime := []int{2,4,0,8,10}

	for i:= 0; i<len(trainName);i++{
		currentTrain := trainName[i]
		currentTrainTime := trainDelaytime[i]

		announcement := train.GetTrainScheduleUpdate(currentTrain,currentTrainTime)

		fmt.Println("Current Trian and It's time",i+1,announcement)
	}


	// fmt.Println("Hello Go!")

	// test.Test()
	// test.VariableExample()
	// test.ArrayPractice()
	// test.SlicePractice()
	// fmt.Println(test.PracticeFunction(1,3,3,4,5,6,6))
	// fmt.Println(test.PracticeFunction(1,2,2,2,2,4))
	// result := test.PracticeClosure()
	// fmt.Println("closure")
	// fmt.Println(result())
	// fmt.Println(result())
	// fmt.Println(result())
	// fmt.Println(result())

	// result2 := test.PracticeClosure()
	// fmt.Println("second Closure")
	// fmt.Println(result2())
	// fmt.Println(result2())
	// fmt.Println(result2())
	// fmt.Println(result2())
	// fmt.Println(result2())
	// fmt.Println("First closure Value Check:")
	// fmt.Println(result())
	test.SafeDrop()
}


