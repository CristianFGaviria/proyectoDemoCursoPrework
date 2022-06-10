package main

import "fmt"

func main() {
	//Declaracion de constates
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi1 =", pi)
	fmt.Println("pi2 =", pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El Area cuadrado es: =", areaCuadrado)

}
