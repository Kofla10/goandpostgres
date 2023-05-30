package main

import (
	"database/sql"
	_ "database/sql" //This package enable us create querys sql
	"log"

	_ "github.com/lib/pq" // Este paquete nos permite realizar la conexion a la bd
	// _ el guion bajo se pone para ingnorar el paquete que estamos llamando si no se va a utilizar
)

// getConnection obtiene una conexión a la base de datos
func getConnection() *sql.DB {
	//creamos la cadena de la conexión
	// dsn := "postgres://'db user':'pss user'@'direction service':'port'?'sslmode'='enabel or disable'" el sslmode quiere decir para realizar la conexión segura a la base de datos
	dsn := "postgres://golang:golang@120.0.0.1:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn) //abrimos una conexion a la base de datos, la funcion open nos devuelve un error y una conexion a la basede datos

	if err != nil {
		// fmt.Println("Error en la conexión a la base de datos")
		log.Fatalf("Error al conectarse a la base de datos => %s", err)
	}

	return db
}

/**
Comando para crear un usuario en go
create user "name user" password "pass user"

creamos base de datos y le damos cual es el usuario dueño, para este ejemplo la base de datos se llamara gocrud
create database gocrud owner "name users"

creacion de la tabla estudiantes

create table estudiantes (
	id serial not null,
	name varchar(50) not null,
	age smallint not null,
	active boolean not null,
	created_at timestamp not nuyll default now(),
	updated_at timestamp);
)
*/
