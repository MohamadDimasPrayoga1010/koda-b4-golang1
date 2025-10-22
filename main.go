package main

import (
	"fmt"
	"minitask1.go/circle"
)

// func areaOfCircle(r float64) float64{
// 	const phi = 3.14
// 	return phi * r * r
// }

// func circumference(r float64)float64{
// 	const phi = 3.14
// 	return 2 * phi * r
// }

func main(){	
	var r float64
	fmt.Println("Masukan Nilai Untuk Menghitung Luas & Keliling Lingkaran")
	fmt.Scan(&r)

	c := circle.Circle{Radius: r}

	fmt.Printf("Jari-jari: %.2f\n", c.Radius)
	fmt.Printf("Luas lingkaran: %.2f\n", c.Area())
	fmt.Printf("Keliling lingkaran: %.2f\n", c.Circumference())


	// fmt.Println("Hasil Luas Keliling Lingkaran",r ,"adalah", areaOfCircle(r))
	// fmt.Println("Hasil Keliling Lingkaran",r, "adalah",circumference(r))
}