package main

import "fmt"

type murid struct {
	name string
	class int
}

func main(){
	var s1 murid
	s1.name = "Aldo"
	s1.class = 3

	fmt.Println("Nama Saya",s1.name)
	fmt.Println("Kelas Saya",s1.class)

}