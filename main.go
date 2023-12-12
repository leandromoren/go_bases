package main

import (
	"fmt"

	"github.com/leandromoren/go_bases.git/variables"
)

func main() {
	//variables.MostrarVariablesEnteras()
	//variables.RestoVariables()
	estado, texto := variables.ConvierteATexto(50)
	fmt.Println(estado, texto)
}
