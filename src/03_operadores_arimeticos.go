package main

import "fmt"

func main() {
	x := 10
	y := 50

	// Suma

	result := x + y
	fmt.Println("Suma:", result)

	// Resta

	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incrementar
	x++
	fmt.Println("Incremetal:", x)

	// Decrementar
	x--
	fmt.Println("Decremental:", x)

	// Reto
	// area del Rectangulo, Trapecio y circulo

	// Rectangulo
	base := 4
	altura := 2

	println("Area del Rectangulo:", base*altura)

	// Trapecio
	base = 3
	baseAbajo := 2
	areaTrapecio := altura * (base + baseAbajo) / 2
	println("Area del Trapecio:", areaTrapecio)

	//Circulo
	const pi = 3.14
	const radio = 6
	areaCirculo := pi * (radio * radio) / 2
	println("Area del circulo:", areaCirculo)
}
