package main

import "fmt"

func main() {
	myArray := [10]int{1,2,3,4,5,6,7,8,9,10}

	var mySlice []int = myArray[:5]
	fmt.Println("array is")
	for _,v := range myArray{
		fmt.Print(v ," ")
	}
	fmt.Println("\n slice is")
	for _,v := range mySlice{
		fmt.Print(v ," ")
	}
	fmt.Println("\n@@@@@@@@@@")

	mySlice1 := make([]int,5)
	mySlice2 :=make([]int,5,10)
	mySlice3 :=[]int{1,2,3,4,5}
	print(mySlice1)
	print(mySlice2)
	print(mySlice3)

}

func print(slice []int){
	fmt.Println("\n---------")
	for _,v := range slice{
		fmt.Print(v ," ")
	}
}
