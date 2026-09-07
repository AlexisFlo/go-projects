package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaración de constantes
const Pi float32 = 3.14

const (
	x = 100
	y = 0b100 // binario
	z = 0o12  // Octal
	w = 0xFF  // Hexadecimal
)

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("Hola")
	fmt.Println(quote.Hello())

	// Variables
	// dentro de las funciones se usa := para declarar variables
	Firstname, lastName, age := "Mochi", "Morchi", 27
	fmt.Println("Te amo", Firstname, lastName, age)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Friday)
}
