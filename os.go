package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("Argument :",args)
	// fmt.Println(args[1])
	// fmt.Println(args[2])

	name,err := os.Hostname()

	if err == nil{
		fmt.Println("Hostname : ",name)
	} else {
		fmt.Println("Not Found")
	}
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	// coba lakukan ini di konsol
	// export APP_USERNAME
	// export APP_PASSWORD
}