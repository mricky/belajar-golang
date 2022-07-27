package main

import "fmt"

func Ups(i int) interface{} {
	if i==1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}
	// function ini bisa mengembalikan number, bolean, atau string
	// karna kontrak interface tidak ekplisit harus mengeluarkan output apa
func main(){
	//var data int = Ups(1) ini tidak bisa dilakukan
	message := Ups(10)
	fmt.Println(message)
}