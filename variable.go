package main

import "fmt"

func main(){
	var name string;
	
	name = "Mohammad Ricky";
	fmt.Println(name);

	var friendName = "Ucil"; // tampa declarasi tipe data
	fmt.Println(friendName) ;

	var age = 30;
	fmt.Println(age);

	country := "Malaysia"; // deklarasi awal, selanjutnya hanya pakai = jika tidak mau pakai var
	country = "Indonesia";
	fmt.Println(country);
	// Deklarasi Multiple Variable
	var (
		firstName = "Rinrin"
		lastName = "Sariningsih"
	)
	fmt.Println(firstName);
	fmt.Println(lastName);

	var (
		productCode = "0001"
		productName = "ProductName"
	)
	fmt.Println(productCode);
	fmt.Println(productName);
}