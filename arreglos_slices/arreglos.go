package arreglos_slices

import "fmt"

var tabla [10]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 10

	fmt.Println(tabla)
}
