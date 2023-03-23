package main

import (
	"fmt"

	pk "./mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Toyota"
	myCar.Year = 2023
	fmt.Println(myCar)

	// Acceso privado
	var myOtherCar pk.CarPrivate
	fmt.Println(myOtherCar)

	pk.PrintMessage("Hola mundo")

	// Si pones una funcion con minuscula al inicio no va a funcionar nunca
}
