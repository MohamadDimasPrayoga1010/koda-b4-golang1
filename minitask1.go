package main

import "fmt"


func areaOfCircle(r float64) float64{
	const phi = 3.14
	return phi * r * r
}

func circumference(r float64)float64{
	const phi = 3.14
	return 2 * phi * r
}

func main(){	
	fmt.Println("Hasil Luas Keliling Lingkaran r = 7", areaOfCircle(7))
	fmt.Println("Hasil Keliling Lingkaran r = 7", circumference(7))
}