package main

import "fmt"

func random() interface{} {
	return "Ricky"
}
func main(){
	var result interface{} = random()

	var resultString string = result.(string)
	fmt.Println(resultString)

	// var resultInt int = result.(int) // panic
	// fmt.Println(resultInt)

	// menggunakan switch
	switch value:= result.(type){
	case string:
		fmt.Println("String",value)
	case int:
		fmt.Println("Int",value)
	default:
		fmt.Println("Unknown")
	}
}