package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string){
	// parameter output nya dapat diubah
	firstName = "Mohammad"
	middleName = "Ricky"
	lastName = "Susanto"

	// optional menyebutkan param return, cukup hanya return saja dapat mengembalikan smua param
	return

}
func main(){
	
	a,b,c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}