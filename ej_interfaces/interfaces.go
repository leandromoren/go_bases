package ej_interfaces

import (
	"fmt"

	"github.com/leandromoren/go_bases.git/interfaces"
)

func HumanosRespirando(hum interfaces.Humano) {
	hum.Respirar()
	fmt.Printf("Soy un %s, y estoy respirando \n", hum.Sexo())
}
