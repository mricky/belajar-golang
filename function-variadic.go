package main

import "fmt"

func sumAll(numbers...int) int{
	total := 0
	for _, value := range numbers{
		total += value
	}
	return total
}
func main(){
	// variadic jika param nya tidak diisi maka akan diisi 0
	// parameter dari fungsi ini menjadi slice []int
	// bisa langsung diisi slice dengan tanda ... di belakang

	total := sumAll(10,2,3,2,3,2,3,)
	slice := []int{1,2,3,4,5,6,7}
	totalSlice := sumAll(slice...)

	fmt.Println(total)
	fmt.Println(totalSlice)
}