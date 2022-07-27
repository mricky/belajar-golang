package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHai(name string){
	fmt.Println("Hello",name, "My Name is",customer.Name)
}
func (customer Customer) sayHuu(name string){
	fmt.Println("Huuuuu uuu ", customer.Name)
}

func main(){
	// Membuat data struct
	var customer  Customer
	customer.Name = "Ricky"
	customer.Address = "Bandung"
	customer.Age = 34
	customer.sayHai("Joko"); // seakan2 struct punya function, padahal bukan dan berdisi sendiri
	customer.sayHuu("Joko")

	fmt.Println(customer)

	// Struct Literal, cara bikin struct
	rinrin := Customer{
		Name: "Rinrin",
		Address: "Bandung",
		Age: 33,
	}

	gavyn := Customer{"Gavin","Bandung",10}
	fmt.Println(rinrin)
	fmt.Println(gavyn)

	// Struct Method
}