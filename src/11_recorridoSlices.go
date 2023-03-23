package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string
	/* Reto: validar que si el texto contiene
	mayúsculas igual siga validando el palíndromo. */
	text = strings.ToLower(text)

	// Eliminamos los espacios
	text = strings.ReplaceAll(" ", text, text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	// Saltear el indice
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// Obtener el indice solo
	for i := range slice {
		fmt.Println(i)
	}

	isPalindrome("Anita lava la tina")
}
