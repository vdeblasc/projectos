package ejemplos

import "fmt"

func Arrays() {
	var alumnos [3]string

	fmt.Println("Cual es le nombre del primer alumno?")
	fmt.Scanf("%s", &alumnos[0])

	fmt.Println("Cual es le nombre del segundo alumno?")
	fmt.Scanf("%s", &alumnos[1])

	fmt.Println("Cual es le nombre del tercer alumno?")
	fmt.Scanf("%s", &alumnos[2])

	fmt.Println("El primer alumno es:", alumnos[0])
	fmt.Println("El segundo alumno es:", alumnos[1])
	fmt.Println("El tercer alumno es:", alumnos[2])
}
