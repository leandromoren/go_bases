package main

import (
	"github.com/leandromoren/go_bases.git/middleware"
)

func main() {
	//variables.MostrarVariablesEnteras()
	//variables.RestoVariables()
	/*estado, texto := variables.ConvierteATexto(50)
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
	fmt.Println(texto) */

	//fmt.Println(ejercicios.TablaMultiplicar())

	//files.SumaTabla()
	//files.LeerArchivo()

	//funciones.LlamarClosure()
	//funciones.Exponencia(2)
	//arreglos_slices.MuestroArreglos()
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	// Pedro := new(modelos.Hombre)
	// e.HumanosRespirando(Pedro)
	//defer_panic.EjemploPanic()

	// Rutina asincrona con la palabra <go>
	// go goroutines.MiNombreLento("Leandro Moren")
	// fmt.Println("Estot aca")
	// var x string
	// fmt.Scanln(&x)

	//webserver.MiWebServer()

	middleware.MiMiddleware()
}
