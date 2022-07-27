package main

import "fmt"

func main(){
	// Slice Potongan dari data array
	// Slice mirip dengan array, yang membedakan adalah ukuran slice bisa berubah
	// Slice dan Array selalu terkoneksi, dimana slice data yang mengkases sebagian atau seluruh data di Array
	// Tipe Data Slice memiliki 3 data, pointer, length, dan capacity
	// Pointer adalah penunjuk data pertama di array per slice
	// Length adalah panjang dari slice, dan
	// Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity
	// array(low:high) // dimulai dari index low sampai sebelum high
	// array[low:] // dimulai dari index low sampai index akhir
	// array[:high] // dimulai index 0 sampai index sebelum sebelum high
	// array[:] // dimulai index 0 sampai index akhir array

	// ... jika ga tau jumlah kapasitas nya

	// Function Slice
	// len(slice)
	// cap(slice)
	// append(slice)
	// make([]TypeData, length, capacity)
	// copy(destination, source)

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
	fmt.Println(cap(slice)) // jumlah kapasitas
	

	// append
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Ricky") // membuat slice baru
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