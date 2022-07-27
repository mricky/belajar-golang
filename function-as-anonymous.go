package main

import "fmt"

type Blacklist func(string)bool // buat alias

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You are blocked",name)

	} else {
		fmt.Println("Welcome "+name)
	}
}
func main(){
	blacklist := func(name string)bool{ // tanpa nama function 
		return name == "admin"
	}
	registerUser("admin",blacklist)
	registerUser("root",blacklist)

	registerUser("operator", func(name string)bool{
		return name=="operator"
	})

	


}