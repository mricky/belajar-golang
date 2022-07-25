package main

import "fmt"

func getFullName() (string, string,string){
	return "Mohammad", "Ricky","Susanto"
}
func main(){
	
	firstName, _,lastName := getFullName()

	fmt.Println(firstName)
	// middlename misal di ignore, gunakan tanda _ garis bawah
	//fmt.Println(middlename)
	fmt.Println(lastName)
}