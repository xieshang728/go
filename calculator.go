package main

import "fmt"

type SlaryCalculator interface {
	CalculateSalary() int
}

type Contract struct{
	empId int
	basicpay int
}

type Permanent struct{
	empId int
	basicpay int
	jj int
}

func (p Permanent) CalculateSalary() int{
	return p.basicpay + p.jj
}

func (c Contract) CalculateSalary() int{
	return c.basicpay
}

func total(slarys []SlaryCalculator) int{
	expense := 0
	for _,s := range slarys{
		expense = expense + s.CalculateSalary()
	}
	return expense
}

func main(){
	emp1 := Permanent{1,300,200}
	emp2 := Permanent{2,200,400}
	emp3 := Contract{1,100}

	employees := []SlaryCalculator{emp1,emp2,emp3}

	expense := total(employees)
	fmt.Printf("expense %d ",expense)

}