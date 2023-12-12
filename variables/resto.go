package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Leandro"
	Estado = true
	Sueldo = 240000.00
	Fecha = time.Now().Local()

	fmt.Println("========================================")
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
	fmt.Println("========================================")
}

// Le paso un numero y retorna un bool y un string
func ConvierteATexto(numero int) (bool, string) {
	//Convierte entero a alfanumerico
	texto := strconv.Itoa(numero)
	return true, texto
}
