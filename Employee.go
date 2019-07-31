package main

import (
	"time"
)

type Employee struct{
	Id int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerId int
}