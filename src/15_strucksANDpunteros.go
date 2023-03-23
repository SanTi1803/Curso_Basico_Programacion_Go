package main

import "fmt"

type PC struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc PC) ping() {
	fmt.Println(myPc.Brand, "Pong")
}

func (myPc *PC) DuplicateRAM() {
	myPc.Ram *= 2

}

func main() {
	a := 50
	// El “&” accede a la dirección del espacio de memoria de la variable.
	b := &a

	/* “*” accede al valor que contiene ese espacio de memoria, dado el nombre de una
	variable o una dirección especifica.*/
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := PC{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.DuplicateRAM()

	fmt.Println(myPc)
	myPc.DuplicateRAM()

	fmt.Println(myPc)
}
