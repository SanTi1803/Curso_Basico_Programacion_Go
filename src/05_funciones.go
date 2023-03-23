package main

import (
	"fmt"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleRetrun(a int) (c, d int) {
	return a, a * 2
}

// Reto
// Area de Rectangulo, Perimetro y circulo con funciones

func calcularAreaRectangulo(base, altura int) {
	areaRectangulo := base * altura
	fmt.Println("Area del Rectangulo:", areaRectangulo)
}

func calcularAreaTrapecio(a, b1, b2 float64) {
	areaTrapecio := a * (b1 + b2) / 2
	fmt.Println("Area del Trapecio:", areaTrapecio)
}

func calcularAreaCirculo(r float64) {
	const pi float64 = 3.14
	areaCirculo := pi * (r * r)
	fmt.Println("Area del circulo:", areaCirculo)
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	//value1, value2 := doubleRetrun(2)
	//fmt.Println("Value1 y Value2:", value1, value2)

	value1, _ := doubleRetrun(2)
	fmt.Println("Value1", value1)

	calcularAreaRectangulo(5, 6)
	calcularAreaTrapecio(15, 12, 6)
	calcularAreaCirculo(10)

}
