package main

import "fmt"

func main() {
	v, err := checkMark()

	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(v)
}

func checkMark() (result string, err error) {
	// Declaramos la variable donde vamos a almacenar la nota leida por pantalla
	var mark int
	// Leemos input
	_, err = fmt.Scanf("%d", &mark)

	// En caso de error, no devolvemos resultado, sino el error
	if err != nil {
		return "", err
	}

	// Manejamos casos
	switch {
	case mark >= 0 && mark < 5:
		result = "Insuficiente"
	case mark >= 5 && mark < 7:
		result = "Suficiente"
	case mark >= 7 && mark < 9:
		result = "Notable"
	case mark >= 9 && mark < 11:
		result = "Sobresaliente"
	default:
		result = "Nota no válida"
	}
	// En caso de exito, el error vale nil y devolvemos resultado
	return result, nil
}
