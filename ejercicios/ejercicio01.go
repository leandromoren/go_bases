package ejercicios

import (
	"strconv"
)

func EjercicioUno(cadena string) (int, string) {
	convertCadena, error := strconv.Atoi(cadena)
	if error != nil {
		return 0, "Hubo un error " + error.Error()
	}
	if convertCadena > 100 {
		return convertCadena, "Es mayor a 100"
	} else {
		return convertCadena, "Es menor a 100"
	}
}
