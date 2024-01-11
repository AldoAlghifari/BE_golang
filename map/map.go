package main

import "fmt"

func main(){
	var angka map[string] int
	angka = map[string] int{}
	angka["number"] = 40
	angka["huruf"] = 60
	fmt.Println(angka["number"])
}