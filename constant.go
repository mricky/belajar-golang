package main

import "fmt"

func main(){
	// variable yang tidak bisa diubah
	const fistName string = "Mohammad"
	const lastName = "Ricky"
	const value = 1000

	fmt.Println(fistName)
	fmt.Println(lastName)

	// deklarasi multiple constan
	const (
		alamat = "Banjar"
		sekolah = "Politeknikpos"
	)
	fmt.Println(alamat)
	fmt.Printf(sekolah)

}