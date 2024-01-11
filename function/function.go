package main

import (
	"fmt"
	"strings"
)

func main(){
	var name = []string{"Aldo", "Alghifari"}
	printMessage("Hai", name)
}

func printMessage(message string, arr[]string){
	var nameString = strings.Join(arr, " ")

	fmt.Println(message, nameString)
}