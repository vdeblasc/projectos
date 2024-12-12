package ejemplos

import "fmt"

func sumar(a int, b int) int {
	return a + b
}

func mainEnteros() {
	fmt.Println("La suma de 5 + 6 es = ", sumar(6, 5))
}
