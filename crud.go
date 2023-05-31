package main

import (
	"fmt"
	"log"
)

func main() {
	/** Comentamos el insert a la tabla de estudiantes
	e := Estudiantes{
		// Name: "Erika",
		// Age:    23,
		Active: true,
	}
	err := Crear(e)
	if err != nil {

		log.Fatalf("Error final => %s", err)
	}

	fmt.Println("Estudiante creado exitosamente")
	*/

	/** Traemos la funcion que nos trae los datos de la tabla estudiatnes
	estudiantes, err := Consultar()

	if err != nil {
		fmt.Printf("error en la consulta de los usuarios main => %s", err)
	}
	fmt.Println(estudiantes)
	*/

	/** Traemos la funcion para la actualizacion de un estudiante
	eu := Estudiantes{
		ID:     4,
		Name:   "Gineth",
		Age:    29,
		Active: false,
	}
	err := Actualizar(eu)

	if err != nil {
		log.Fatalf("No se pudo actualizar el registo => %s", err.Error())
	}
	fmt.Println("Estudiantes actualizado con exito")
	**/
	err := Borrar(1)

	if err != nil {
		log.Fatal("Error en la eliminación del estudiante")
	}
	fmt.Println("Estudiante eliminado con Exito")
}
