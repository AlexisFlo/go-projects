package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola")
	fmt.Println(quote.Hello())

	// Variables
	// dentro de las funciones se usa := para declarar variables
	name, lastName, age := "Mochi", "Morchi", 27
	fmt.Println("Te amo", name, lastName, age)
}
