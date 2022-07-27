package main

import "fmt"

func main(){

	// Function Map
	// len(map)
	// map[key]
	// map[key] = value
	// make(map[TypeKey]TypeValue)
	// delete(map,key)

	person := map[string]string{
		"name" : "Mohammad",
		"address" : "Bandung",
	}
	// jika ingin menambahkan key value setelah diinisiasi

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	// Beberapa fungsi pada map yang sering dipakai
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Ricky"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book,"ups")
	fmt.Println(book)
	fmt.Println(len(book))

	user := map[string]string{
		"firstName" : "Mohammad",
		"lastName" : "Ricky",
	}
	delete(user, "lastName");
	fmt.Println(user["firstName"]);

}