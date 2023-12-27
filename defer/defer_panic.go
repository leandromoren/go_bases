package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un msj")
	defer fmt.Println("Este es el msj final")

	fmt.Println("Este es el segundo msj")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		// Sirve para abortar el sistema con un mensaje
		panic("Se encontro el valor 1")
	}
}
