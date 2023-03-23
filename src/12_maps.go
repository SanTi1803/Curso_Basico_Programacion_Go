package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Santi"] = 20

	fmt.Println(m)

	/* Otra manera de hacer un diccionario

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	*/

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["Santi"]
	fmt.Println(value, ok)
}
