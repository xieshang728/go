package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func eval(a, b int, op string) int{
	fmt.Println("----------------")
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("undefined operator: "+op)
	}
	return result
}

func circle(value int) int{
	var result int
	for num := 0; num < value; num ++{
		result += num
	}
	return result
}

func printFile(filename string){
	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main()  {
	const filename = "abc"
	data,error := ioutil.ReadFile(filename)
	if error != nil{
		fmt.Printf("%s------",error)
		return
	}
	fmt.Println(data)
	result := eval(4,2,"/")
	fmt.Printf("result: %d",result)
	fmt.Println("circle method ----------------------")
	fmt.Printf("result %d",circle(10))
	fmt.Println("---------printFile-----------")
	printFile("abc")
}
