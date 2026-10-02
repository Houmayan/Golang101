package test

import "fmt"

func ArrayPractice() {
	arr := [5]int{1, 2, 3, 4, 5}
	// for i:=0;i< 5;i++{
	// 	// fmt.Println(arr[i])
	// }
	for _,val := range arr{
		fmt.Println(val)
	}
	fmt.Println(arr)
}


func SlicePractice(){

	sl:= make([]int,4,6)
	
	fmt.Println(".................\n")
	fmt.Println("Printing Empty Slice",sl)
	fmt.Println("Printing SLice Capacity and Length",len(sl),cap(sl))

	sl = append(sl,5,6) 

	
	fmt.Println("Printing SLice After Appending",len(sl),cap(sl),sl)
	sl = append(sl, 7,8)
	fmt.Println("Printing SLice After Appending twice",len(sl),cap(sl),sl)
	for idx,value := range sl{
		fmt.Println("Slice Range:",idx,value)
	}

}