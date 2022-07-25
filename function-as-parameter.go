package main

import "fmt"

// supaya pemanggilan function nya tidak panjang, dapat mengunakan type declaration / alias

type Filter func(string)string
func sayHalloWithFilter(name string, filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello "+nameFiltered)
}
func spamFilter(name string)string{
	if name == "Anjing"{
		return "....."
	} else 
	{
		return name
	}
}
func main(){
	sayHalloWithFilter("Ricky", spamFilter)
	filter := spamFilter
	sayHalloWithFilter("Anjing",filter)
}