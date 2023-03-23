package main

import (
	"fmt"

	pk "./mypackage"
)

func main() {
	myCuadrado := pk.Cuadrado{Base: 4}
	fmt.Println(myCuadrado)
	pk.Calcular(myCuadrado)

	myRectangulo := pk.Rectangulo{Base: 2, Altura: 4}
	fmt.Println(myRectangulo)
	pk.Calcular(myRectangulo)
}
