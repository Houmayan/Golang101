package test

func PracticeClosure () func() int{

	conter :=0

	return func() int {
		conter++
		return conter
	}

	

}

