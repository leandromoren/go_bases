package middleware

import "fmt"

/*

 */

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	resultado := operacionesMidd(sumar)(5, 5)
	resultado2 := operacionesMidd(restar)(5, 5)
	resultado3 := operacionesMidd(multiplicar)(5, 5)

	fmt.Println(resultado)
	fmt.Println(resultado2)
	fmt.Println(resultado3)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("Inicio de operacion")
		return f(i1, i2)
	}
}
