package main

import (
	"fmt"
	"runtime"

	"github.com/leandromoren/go_bases.git/ejercicios"
	"github.com/leandromoren/go_bases.git/variables"
)

func main() {
	//variables.MostrarVariablesEnteras()
	//variables.RestoVariables()
	estado, texto := variables.ConvierteATexto(50)
	fmt.Println(estado, texto)

	//Puedo declarar la variable dentro del condicional
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("El sistema operativo es: " + os)
	} else if os == "windows" {
		fmt.Println("El sistema operativo es: " + os)
	} else {
		fmt.Println("El sistema operativo es: " + os)
	}

	numero, texto := ejercicios.EjercicioUno("99")
	fmt.Println(numero)
	fmt.Println(texto)

}
