package main

import (
	"fmt"
	"math"
)

func main(){
	area, circum := calculate(3.5)
	fmt.Println(area)
	fmt.Println(circum)
}

func calculate(d float64)  (float64, float64) {
	var area = math.Pi * math.Pow(d/2,2)
	var circumference = math.Pi * d

	return area, circumference
}

