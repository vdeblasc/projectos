package ejemplos

import "fmt"

func hola(s string) string {
	return "Hola " + s
}

func mainArgumento() {
	fmt.Println(hola("Luis"))
}
