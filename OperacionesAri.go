package main

import "fmt"

func main() {
	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("La suma es:", result)
	//resta
	result = y - x
	fmt.Println("La resta es:", result)

	largo := 60
	ancho := 40
	rectangulo := largo * ancho
	fmt.Println("Area rectangulo=", rectangulo)

	a := 20
	b := 30
	c := 50
	trapecio := (a + b) / 2 * c
	fmt.Println("Area Trapecio =", trapecio)

	const pi float64 = 3.14
	var r float64 = 5

	circulo := pi * r * r
	fmt.Println("Area Circulo =", circulo)
}
