package main

import "fmt"

func main(){
	name := "Ricky"

	switch name {
	case "Ricky":
			fmt.Println("Hello Ricky")
			fmt.Println("Hello Ricky") // bisa execute lebih satu baris
	case "Joko" :
			fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Asik..")

	}

	// switch mendukung short expresion

	// switch length := len(name); length > 5 {
	// case true :
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false :
	// 	fmt.Println("Nama Sudah Benar")
	// }

	// switch tanpa expression / kondisi
	
	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Sudah Benar")
	}

}