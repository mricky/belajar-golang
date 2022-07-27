package main

import "fmt"

type Address struct {
	CITY string
	PROVINCE string
	STATE string
}

func ChangeAddressToIndoneisa(address *Address)
func main(){

	// contoh pass by value
	// address1 := Address{"Bandung","West Java","Indonesia"}
	// address2 := address1
	// address2.CITY = "Banjar"
	// fmt.Println(address1)  // output {Bandung West Java Indonesia} address1 tidak berubah
	// fmt.Println(address2) // {Banjar West Java Indonesia}

	//  Operator &  
	// address1 := Address{"Bandung","West Java","Indonesia"}
    // address2 := &address1 // ini akan menjadi me reference ke address1

	// address2.CITY = "Banjar" // merubah field
	// fmt.Println(address1) // {Banjar West Java Indonesia} coba laukan cmd + point ke variable
	// fmt.Println(address2) // &{Banjar West Java Indonesia}

	// Operator *
	// Case Tanpa Operator * 
	//  address1 := Address{"Bandung","West Java","Indonesia"}
    //  address2 := &address1 // ini akan menjadi me reference ke address1

	//  address2.CITY = "Banjar" // merubah field
	//  //address2 = Address{"England","Britis", "Inggris"} // ini tidak bisa dilakukan karna addres sudah menjadi pointer
	//  address2 = &Address{"England","Britis", "Inggris"} 
	//  fmt.Println(address1) // {Banjar West Java Indonesia} 
	//  fmt.Println(address2) // &{England Britis Inggris}  hal ini tidak merubah ke address1, tpi membuat data baru

	// Case dengan Operator *
	// Sample 1
	//  address1 := Address{"Bandung","West Java","Indonesia"}
    //  address2 := &address1 // ini akan menjadi me reference ke address1

	//  address2.CITY = "Banjar" // merubah field
	//  //address2 = Address{"England","Britis", "Inggris"} // ini tidak bisa dilakukan karna addres sudah menjadi pointer
	//  *address2 = Address{"England","Britis", "Inggris"} 
	//  fmt.Println(address1) // {England Britis Inggris}
	//  fmt.Println(address2) // &{England Britis Inggris}  di paksa nge reference ke data yang baru

	// Sample 2
	   var address1 Address = Address{"Bandung","West Java","Indonesia"}
	   var address2 *Address = &address1
	   var address3 *Address = &address1

	   *address2 = Address{"Banjar","West Java","Indonesia"}
	   fmt.Println(address1)
	   fmt.Println(address2)
	   fmt.Println(address3)

	// Function New
	// Membuat pointer dan isi data nya kosong
	// hanya mengembalikan pointer ke data kosong artinya tidak ada data awal

	var address4 = new(Address)
	fmt.Println(address4) // output &{  }
}