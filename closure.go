package main

import "fmt"
func main(){
	name :="Budi";
	counter := 0;

	increment := func(){
		name = "Ricky";
		fmt.Println("Increment");
		counter++;
	}

	increment();
	increment();
	fmt.Println(counter)
	fmt.Println(name)
}