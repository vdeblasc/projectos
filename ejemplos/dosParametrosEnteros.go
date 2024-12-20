package ejemplos

import "fmt"

func Sumar(a int, b int) int {
	return a + b
}

func Enteros() {
	fmt.Println("La suma de 5 + 6 es = ", Sumar(6, 5))
}
