package main

import "fmt"

func sayHello(){
	fmt.Println("Hello World")
	fmt.Println("Hello Dunia")
}
func main(){
	sayHello()
	for i:=0; i< 5; i++{
		sayHello()
	}
}