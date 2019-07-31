package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	username = "xieshang"
	age = 11
)

func variable(){
	var a int
	var s string
	fmt.Printf("%d  %q\n",a,s)
}

func variableInit(){
	var a int = 3
	var s string = "i am a string"
	fmt.Printf("%d %q\n",a,s)
}

func variableDeduction(){
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)

}

func variableShorter(){
	a,b,c,s := 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func euler(){
	c := 3+4i
	fmt.Println(cmplx.Abs(c))
}

func triangle(){
	var a,b = 3,4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts(){
	const a,b = 3,4
	const filename = "abc.txt"
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(c)
	fmt.Println(filename,c)
}
func enums(){
	fmt.Println("---------------------")
	const(
		java = iota
		python
		cpp
		golang
	)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(java,python,cpp,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func bounded(v int) int{
	if v < 100 {
		return 100
	}else if v > 100{
		return 0
	}else{
		return 100
	}
}


func main()  {
	fmt.Println("hello world")
	fmt.Printf("%s---%d",username,age)
	variable()
	variableInit()
	variableDeduction()
	variableShorter()
	euler()
	triangle()
	fmt.Println("-------------------------")
	consts()
	enums()
	fmt.Println("-------------------------")
	result := bounded(100)
	fmt.Printf("result %d",result)
}
