package main

import "fmt"

func main(){
	months := [...]string{1:"January",2:"February",3:"March",4:"April",5:"Mary",
		6:"June",7:"July",8:"August",9:"September",10:"October",11:"November",12:"December"}
	fmt.Println(months)
	fmt.Println(len(months))
	fmt.Println(cap(months))

	fmt.Println()
}
