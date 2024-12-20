package ejemplos

import (
	"fmt"
	"reflect"
)

func Verificar() {
	var persona string = "Jose Luis"
	var edad uint8 = 22
	var peso uint16 = 90

	fmt.Println("La variable persona es de tipo", reflect.TypeOf(persona), "y su valor es", persona)
	fmt.Println("La variable edad es de tipo", reflect.TypeOf(edad), "y su valor es", edad)
	fmt.Println("La variable peso es de tipo", reflect.TypeOf(peso), "y su valor es", peso)

	fmt.Printf("La variable alumno es de tipo %T y su valor es %v \n", persona, persona)
	fmt.Printf("La variable alumno es de tipo %T y su valor es %v \n", edad, edad)
	fmt.Printf("La variable alumno es de tipo %T y su valor es %v \n", peso, peso)
	fmt.Printf("La persona se llama %v, su edad es %v y su peso es %v \n", persona, edad, peso)
}
