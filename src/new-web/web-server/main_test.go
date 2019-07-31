package web

import (
	"fmt"
	"testing"
)

//func TestMain(m *testing.M){
//	fmt.Println("test main")
//	m.Run()
//}

func TestPrint(t *testing.T){
	res := Print1To20()
	fmt.Println("hey")
	if res != 212{
		t.Error("wrong result of print1to20")
	}
}

func TestPrint2(t *testing.T){
	res := Print1To20()
	fmt.Println("hey")
	res ++
	if res != 212{
		t.Error("wrong result of print1to20")
	}
}

func TestAll(t *testing.T){
	t.Run("TestPrint",TestPrint)
	t.Run("TestPrint2",TestPrint2)
}

func Benchmark(t *testing.B){
	for n := 0; n < t.N; n++{
		Print1To20()
	}
}