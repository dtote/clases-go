package main

import (
	"fmt"
	// . "github.com/dtote/clases-go/saludos" --> Con esta linea podriamos usar las funciones sin necesidad de prefijo
	s "github.com/dtote/clases-go/saludos"
	"github.com/logrusorgru/aurora/v3"
)

func main() {
	fmt.Println(aurora.Blue("Estamos haciendo cositas"))
	s.Saludar()
}
