package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int,error){
	*c += ByteCounter(len(p))
	fmt.Printf("count-----%d",*c)
	fmt.Println("------")
	return len(p),nil
}

func main(){
	var c ByteCounter
	c.Write([]byte("hello"))


	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c,"hello,%s",name)
	fmt.Println(c)
}