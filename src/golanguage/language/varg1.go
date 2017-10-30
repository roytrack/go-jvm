package main

import "fmt"

func MyPrintf(args ...interface{}){
	for _,arg := range args{
		switch arg.(type) {
		case int:
			fmt.Println("this is a int ",arg)
		case string:
			fmt.Println("this is a string ",arg)
		case int64:
			fmt.Println("this is a int64 ",arg)
		default:
			fmt.Println("this is a unknown ",arg)
		}
	}
}

func main() {
	var v1 int =1
	var v2 int64 = 55
	var v3 string = "ssds"
	var v4 bool = false
	var v5 float32 = 1.2323
	MyPrintf(v1,v2,v3,v4,v5)
}
