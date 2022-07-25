package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool

	var noKTPRicky NoKTP = "123141231241231231"
	var marriedStatus Married = true

	fmt.Println(noKTPRicky)
	fmt.Println(marriedStatus)
}