package ejemplos

import "fmt"

func Hola(s string) string {
	return "Hola " + s
}

func Argumento() {
	fmt.Println(Hola("Luis"))
}
