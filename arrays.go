package main

import "fmt"

func main(){
	var arr1 [5]int
	arr2 := [3]int {1,2,3}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int
	fmt.Println(arr1,arr2,arr3)
	fmt.Println("-----------array-------------")
	fmt.Println(grid)
	fmt.Println(grid[1][2])
	fmt.Println("for array---------------")
	for i:= 0; i<len(arr3); i++{
		fmt.Println(arr3[i])
	}
	fmt.Println("------------array2----------")
	for i,value := range arr3 {
		fmt.Printf("index:%d------>value:%d",i,value)
		fmt.Println()
	}
	fmt.Print("---------------------------")
}
