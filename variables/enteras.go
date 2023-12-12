package variables

import (
	"fmt"
)

// Para que una variable pueda ser accedida desde fuera, debe empezar en mayusculas
func MostrarVariablesEnteras() {
	//Manera declarativa
	var intBasico int

	//Manera de asignacion
	//- Entero de 32 bits y 64
	intDe32 := int32(10)
	intDe64 := int64(100)

	fmt.Println("intBasico = ", intBasico)
	fmt.Println("intDe32 = ", intDe32)
	fmt.Println("intDe64 = ", intDe64)
}
