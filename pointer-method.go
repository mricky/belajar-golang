package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main(){
	ricky := Man{"Ricky"}
	ricky.Married()
	fmt.Println(ricky.Name) // Mr. Ricky
}