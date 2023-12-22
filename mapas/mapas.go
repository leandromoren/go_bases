package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)
	// fmt.Println(paises)

	paises["Argentina"] = "Buenos Aires"
	paises["Alemania"] = "Berlin"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	//Otra forma de crear un mapa
	capitales := map[string]int{
		"Lisboa":     5,
		"Washington": 10,
		"Brasilia":   15,
		"Madrid":     20,
	}
	fmt.Println(capitales)

	//Recorrer toda la coleccion
	for capital, puntaje := range capitales {
		fmt.Printf("Capital: %s, tiene un puntaje de %d \n", capital, puntaje)
	}

	//Borrar elemento del mapa
	delete(capitales, "Madrid")
	fmt.Println(capitales)
}
