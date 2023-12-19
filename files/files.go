package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leandromoren/go_bases.git/ejercicios"
)

var pathTxt string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var textoAGrabar string = ejercicios.TablaMultiplicar()

	archivo, err := os.Create(pathTxt)
	if err != nil {
		fmt.Println("Error al crear el archivo ", err.Error())
		return
	}

	fmt.Fprintln(archivo, textoAGrabar)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(pathTxt, texto) {
		fmt.Println("Error al concatenar contenido.")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(pathTxt, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append ", err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString ", err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeerArchivo() {
	archivo, err := os.Open(pathTxt)
	if err != nil {
		fmt.Println("Error al leer archivo ", err)
		return
	}

	scanner := bufio.NewScanner(archivo)
	var count int = 0

	for scanner.Scan() {
		registro := scanner.Text()
		count++
		fmt.Println(count, "> ", registro)
	}
}
