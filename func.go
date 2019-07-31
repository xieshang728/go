package main

import "fmt"

func div(a, b int) (int,int){
	return a*b,a+b
}

func div2(a,b int) (q,r int)  {
	q = a * b
	r = a +b
	return q,r
}

func sum(numbers... int) int{
	sum := 0
	for i:= range numbers{
		sum += numbers[i]
	}
	return sum
}

func swap(a,b *int){
	*a,*b = *b,*a
}

func main(){
	result1,result2 := div(3,4)
	fmt.Printf("%d---%d",result1,result2)
	fmt.Println("------------------")
	q,_ := div2(3,5)
	fmt.Println(q)
	fmt.Println("------sum--------")
	fmt.Println(sum(10,1,2))
	a,b := 3,4
	swap(&a,&b)
	fmt.Println("-----------swap------------")
	fmt.Printf("%d-------->%d",a,b)
}