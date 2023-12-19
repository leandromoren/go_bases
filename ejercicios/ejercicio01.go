package ejercicios

import (
	"bufio"
	"fmt"
	"os"
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

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var texto string

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}

	return texto
}
