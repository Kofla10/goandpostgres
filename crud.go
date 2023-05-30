package main

import (
	"fmt"
	"log"
)

func main() {
	e := Estudiantes{
		Name:   "Camilo",
		Age:    30,
		Active: true,
	}

	err := EstudianteCrear(e)
	if err != nil {
		log.Fatalf("Error final => %s", err)
	}

	fmt.Println("Estudiante creado exitosamente")
}
