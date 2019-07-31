package web

import "fmt"

func Print1To20() int {
	res := 0
	for i := 1; i <= 20; i++{
		res += i
	}
	return res
}

func main(){
	username := "xieshang"
	fmt.Printf("name %s",username)
	fmt.Sprintf()
}
