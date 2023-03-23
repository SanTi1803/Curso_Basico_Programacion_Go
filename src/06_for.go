package main

import "fmt"

func main() {
	// For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For While
	counter := 0
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("\n")

	// For forever

	//counterForever := 0
	//for {
	//	fmt.Println(counterForever)
	//	counterForever++
	//}

	// Reto
	decreter := 10
	for decreter >= 0 {
		fmt.Println(decreter)
		decreter--
	}
}
