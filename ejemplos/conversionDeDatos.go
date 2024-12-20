package ejemplos

import (
	"fmt"
	"strconv"
)

func Conversion() {
	var mayorDeEdad string = "true"
	//:= es otra forma de declarar otra variable sin utilizar var ni string

	mayorDeEdadBool, _ := strconv.ParseBool(mayorDeEdad)

	//%v = valor de la variable
	//%T = tipo de la variable
	fmt.Printf("MayorDeEdad es %v y su tipo es %T\n", mayorDeEdad, mayorDeEdad)
	fmt.Printf("MayorDeEdadBool es %v y su tipo es %T\n", mayorDeEdadBool, mayorDeEdadBool)

	mayorDeEdadStr := strconv.FormatBool(false)

	fmt.Printf("MayorDeEdadStr es %v y su tipo %T", mayorDeEdadStr, mayorDeEdadStr)

	//fmt.Println("Es mayor de edad?", mayorDeEdad)

}
