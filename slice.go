package main

import "fmt"

func main(){
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// di slice terdapat, index, length, capacity
	var slice = months[4:7]
	fmt.Println(slice)

	// function slice
	fmt.Println(len(slice)) // jumlah slice
	fmt.Println(cap(slice)) // julah kapasitas
	
	// append
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Ricky")
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)
 
	// rekomedasi membuat slice yang paling aman
	newSlice := make([]string, 2,5) // 2 (length), 5 (capacity)

	newSlice[0] = "Mohammad"
	newSlice[1] = "Ricky"

	fmt.Println(newSlice)

	// script copy slice
	copySlice := make([]string,2,5)
	copy(copySlice, newSlice)  // copy(destinantion, source)
	fmt.Println(copySlice)

	// hati2 saat membuat array, array harus didefinisikan length nya, klo pun tidak tambahkan ...
	iniArray := [...]int{
		1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	
}