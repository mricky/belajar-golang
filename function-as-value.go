package main

import "fmt"

func getGoodBy(name string) string {
	return "Good Bye " +name
}
func main(){
	goodbye := getGoodBy
	
	result := goodbye("Ricky")
	fmt.Println(result)
}