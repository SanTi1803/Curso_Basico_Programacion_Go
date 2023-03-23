package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "world"

	// Println:  Hace un salto de linea al terminar el mensaje
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf:  Imprime con un formato
	// Cuando sabemos que tipo de variables son podemos usar
	/*
	 bool:                    %t
	 int, int8 etc.:          %d
	 uint, uint8 etc.:        %d, %#x if printed with %#v
	 float32, complex64, etc: %g
	 string:                  %s
	 chan:                    %p
	 pointer:                 %p
	*/

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)

	// Si desconocemos el tipo de variable en el formato
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf:  Genera un string pero no lo imprime en consola
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	// Usamos %T para determinar el tipo de dato de una variable a utilizar.
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
