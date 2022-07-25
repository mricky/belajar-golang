package main

import "fmt"

func main(){
	 name := "Ricky"

	if name == "Ricky" {
		fmt.Println("Hello Ricky")
	} else if name == "Mohammad" {
		fmt.Println("Hello Mohammad")
	} else {
		fmt.Println("Hi, kenalan donk")

	}

	var length = len(name)
	if length > 4 {
		fmt.Println("Terlalu Panjang")
	} else
	{
		fmt.Println("Sudah Benar")
	}
}