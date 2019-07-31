package main

import "fmt"

func main(){
	m := map[string] string{
		"name":"xieshang",
		"course":"math",
		"age":"22",
		"address":"beijing",
	}

	fmt.Println(m)
	fmt.Println("-----------------------------")
	m2 := make(map[string]int)
	fmt.Println(m2)
	fmt.Println("----------------------")

	for k,v := range m {
		fmt.Println(k,v)
	}
	fmt.Println("-----------------------")
	name,ok := m["name"]
	fmt.Println("---------name: ==="+name,ok)
	fmt.Println()

	age,ok := m["age2"]
	fmt.Println("-----------:age,ok"+age,ok)
	m["birthday"] = "1995-07-28"
	fmt.Println(m)

}
