package main

import "fmt"

func main(){
	var names [3]string
	names[0] = "Mohammad"
	names[1] = "Ricky"
	names[2] = "Susanto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array secara langsung saat deklarasi variable

	var values  = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// fungsi di array
	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}