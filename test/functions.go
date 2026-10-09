package test

func PracticeFunction(num1 ...int) int {
	var sum int = 0
	for i := 0; i <= len(num1)-1; i++ {
		sum = sum + num1[i]
	}

	return sum
}

func PracticeFunction2(num2 ...int) (sum2 int) {

	// for i := 0; i < len(num2); i++ {
	// 	sum2 = sum2 + num2[i]
	// }
	// return

	for _, j := range num2 { // range over Function
		sum2 += j
	}
	return
}

