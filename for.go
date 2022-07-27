package main

import "fmt"

func main(){
	
	for counter :=1; counter <= 18; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for range untuk data collection slice / map / array
	
	slice := []string{"Mohammad","Ricky","Susanto"}

	for i:= 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
	for i, value := range slice{
		fmt.Println("Index", i, "=",value)
	}


	// di golang variable harus digunakan, spya variable dapat tidak digunakan, pakai tanda _
	// untuk memberitahu si golang variable i tidak digunakan, so gunakan _
	for _, value := range slice{
		fmt.Println(value)
	}

	person := make(map[string]string)
		person["name"] = "Ricky"
		person["title"] = "Programmer"

	
	for key,value := range person {
		fmt.Println(key,"=",value)
	}
}