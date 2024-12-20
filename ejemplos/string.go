package ejemplos

import "fmt"

func String() {
	var nombre string
	var peso float32

	fmt.Println("Cual es tu nombre?")
	fmt.Scanf("%s", &nombre)

	fmt.Println("Cual es tu peso?")
	fmt.Scanf("%f", &peso)

	fmt.Println("El nombre es: ", nombre)
	fmt.Println("El peso es: ", peso, "kg")
}
