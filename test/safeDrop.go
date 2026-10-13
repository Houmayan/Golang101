package test

import (
	"fmt"
	"math/rand"
)

func SafeDrop() {
	id := rand.Intn(5)
	fmt.Println("Random ID:", id)
	shipment := []string{"car","Wasching Machine","laptop"}
	// prices := []float64 {12,14,143}

	for _, val := range shipment{
		fmt.Println(val)
	}
}