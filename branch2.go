package main

import (
	"bytes"
	"fmt"
	"strings"
)

func basename(filename string) string{
	for i := len(filename)-1; i>= 0; i-- {
		if filename[i] == '/'{
			filename = filename[i+1:]
			break
		}
	}

	for i := len(filename)-1; i >=0 ; i--{
		if filename[i] == '.'{
			filename = filename[:i]
			break;
		}
	}
	return filename
}

func easyBasename(filename string) string{
	index := strings.LastIndex(filename,"/")
	filename = filename[index+1:]
	if dot := strings.LastIndex(filename,"."); dot >= 0{
		filename = filename[:dot]
	}
	return filename
}

func intsToString(values[] int) string{
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i,v := range values{
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf,"%d",v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func intsToChar(){
	x := 12
	y := fmt.Sprintf("%s",x)
	fmt.Printf("%d----->%s",x,y)
}


func main(){
	filename := basename("/a/b/d/c.go")
	fmt.Printf("%s -------filename",filename)
	fmt.Println()
	fmt.Printf("%s--------easyFilename",easyBasename("/a/b/d/c.go"))
	fmt.Println("------------intsToString------------")
	fmt.Println(intsToString([]int{1,2,3}))
	fmt.Println("------------intsToString---------")
	intsToChar()
}