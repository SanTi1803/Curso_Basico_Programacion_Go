package main

import (
	"fmt"

	pk "./mypackage"
)

func main() {
	var componentes pk.PC
	componentes.Brand = "MSI"
	componentes.Disk = 500
	componentes.Ram = 16

	fmt.Println(componentes)
	componentes.DuplicateRAM()
	fmt.Println(componentes)
	componentes.DuplicateRAM()
	fmt.Println(componentes)
}
