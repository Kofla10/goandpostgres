package db

import (
	_ "database/sql" //This package enable us create querys sql

	_ "github.com/lib/pq" // Este paquete nos permite realizar la conexion a la bd
	// _ el guion bajo se pone para ingnorar el paquete que estamos llamando si no se va a utilizar
)
