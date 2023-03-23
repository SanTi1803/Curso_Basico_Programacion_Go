package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad, AND")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	// Reto: numero par e impar
	if pairNumber(6) {
		fmt.Println("Es numero par")
	} else {
		fmt.Println("No es numero par")
	}

	// Reto: Entrar con usuario y contraseña
	if userAndPassword("Santiago", "santifac123") {
		fmt.Println("El ingreso es correcto")
	} else {
		fmt.Println("El usuario o la contraseña estan mal")
	}
}

func userAndPassword(user, password string) bool {
	return user == "Santiago" && password == "santifac123"
}

func pairNumber(n int) bool {
	return n%2 == 0
}
