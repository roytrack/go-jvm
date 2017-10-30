package main

import "fmt"

type PersonInfo struct {
	Id string
	Name string
	Address string
}

func main() {
	var personDb map[string] PersonInfo
	personDb =make(map[string] PersonInfo)
	personDb["555"]=PersonInfo{"555","roy","beijing..."}
	personDb["666"]=PersonInfo{"666","leo","hangzhou..."}
	person,ok :=personDb["555"]
	if ok {
		fmt.Println("found person ",person.Name)
	}else{
		fmt.Println("did not find person ")
	}



}
