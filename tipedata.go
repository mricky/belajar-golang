package main

import "fmt"

func main(){
	// Boolean
	fmt.Println("Benar =",true);
	fmt.Println("Salah =",false);

	// String (Kumulan Karakter)
	fmt.Println("Ricky");
	fmt.Println("Mohammad");
	fmt.Println("Susanto");
	// Function dalam string
	// len("string") jumlah karatket
	// "string"[number"], masukan index karater
	fmt.Println(len("Mohammad Ricky"));
	fmt.Println("Mohamad Ricky"[2]); // output angka adalah represtansikan dari byte

}