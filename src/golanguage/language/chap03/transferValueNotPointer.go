package main

import "fmt"

type Integer int

func (a Integer)Less(b Integer) bool{

	return a<b
}

func (a Integer)Add(b Integer) {
	a+=b
}
func (a *Integer)AddReal(b Integer) {
	*a+=b
}
func main() {
	var  a Integer =10
	a.Add(10)
	fmt.Println(a)
	a.AddReal(10)
	fmt.Println(a)
}
