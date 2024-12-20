package ejemplos

import "fmt"

func ValorCero() {
	var (
		nombre     string
		edad       uint8
		peso       float32
		estudiante bool
	)

	fmt.Printf("nombre:  %v (%T)\n", nombre, nombre)
	fmt.Printf("edad:  %v (%T)\n", edad, edad)
	fmt.Printf("peso:  %v (%T)\n", peso, peso)
	fmt.Printf("estudiante:  %v (%T)\n", estudiante, estudiante)
}
