package main

import "fmt"

func main(){
	var nilai = 6

	if nilai == 10 {
		fmt.Println("Nilai Sempurna")
	} else if nilai >= 7 {
		fmt.Println("Anda lulus")
	} else if nilai <= 6 {
		fmt.Println("Anda Tidak lulus")
	} else if nilai > 2 {
		fmt.Println("Banyak belajar")
	} else {
		fmt.Println("gagal")
	}
}
