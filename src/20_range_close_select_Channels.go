package main

import "fmt"

//len cuantas goroutines (o datos) hay
//Cap cuanta es la cantidad máxima que puede almacenar ese channel

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	// Range y close
	close(c) // sirve para cerrar el canal para no recibir ningun
	// otro dato adicional, lo mejor es cerrar el canal

	//c <- " Mensaje3"
	for message := range c {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)
	go message("Mensaje1", email1)
	go message("Mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
