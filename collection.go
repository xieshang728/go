package main

import "fmt"

func testSlice(){
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0],s[7])
	fmt.Println(s[0:5])
}

func main(){
	fmt.Println("----------testSlice-------------")
	testSlice()
}