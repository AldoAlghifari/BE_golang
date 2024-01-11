package main

import "fmt"

func main(){
	var fruits = []string{"apple", "banana", "grape"}

	
	nameFruits := fruits [:2] //menampilkan data sebelum array ke 2
	// nameFruits := fruits [:] //menampilkan data semua

	fmt.Println(nameFruits)
}