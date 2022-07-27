package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool){
	defer endApp()
	
	if error {
		panic("APLIKASI ERROR")
	}
}
func main(){
	runApp(true)
	runApp(false) // ini tidak di eksekusi
}