package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Mohammad Ricky Susanto","Ricky"))
	fmt.Println(strings.Split("Mohammad Ricky Susanto"," "))
	fmt.Println(strings.ToLower("Mohammad Ricky Susanto"))
	fmt.Println(strings.ToUpper("Mohammad Ricky Susanto"))
	fmt.Println(strings.Trim("  Mohammad Ricky Susanto"   ," "))
	fmt.Println(strings.ReplaceAll("Eko Eko Eko","Eko","Budi"))
}