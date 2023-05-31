package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
)

type Estudiantes struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdateAp  time.Time
}

// crear registro de un studiante
func Crear(e Estudiantes) error {
	//creamos el string para poder insertar en la base de datos
	q := `INSERT INTO 
			estudiantes (name, age, active)
			VALUES ($1, $2, $3)`

	// creacion de valores nulos para la db
	intNull := sql.NullInt64{}
	stringNull := sql.NullString{}

	db := GetConnection() //traemos la conexion a la base de datos
	defer db.Close()      //buena practica es: siempre cerrar la conexión a la base de datos

	//como en golang no se manejan los valores null, entonces realizarmo una validación para que se puedan ingresar valores null a la base de datos
	//Esta validación es para enteros
	if e.Age == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(e.Age)
	}

	//Realizamos la validación para los string
	if e.Name == "" {
		stringNull.Valid = false
	} else {
		stringNull.Valid = true
		stringNull.String = e.Name
	}

	//Realizamos la validación para las fechas

	//creamos la sentencia preparada para que no se pueda realizar la inyecion sql
	stmt, err := db.Prepare(q) //preparamos la sentencia "statement"

	if err != nil {
		log.Fatalf("Error al preparala la sentencia => %s ", err)
		return err
	}

	defer stmt.Close()
	//obtenemos la cantidad de filas que fueron afectadas
	r, err := stmt.Exec(stringNull, intNull, e.Active)
	// el exec se usa para las instrucciones de delete, insert, update, siempre se poner las columnas que se quiere afectar

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}

	return nil
}

// funcion que trea los datos de la tabla estudiantes
func Consultar() (estudiantes []Estudiantes, er error) {
	q := `SELECT 
		  	id, name, age, active, created_at, updated_at
		  FROM estudiantes`

	//creamos el null para las fechas
	//si hay dos variables por ejemplo string null, toca crear variables tipo null para cada valor de la stuctura
	timeNull := pq.NullTime{} //traemos el pq.NullTime de la libreria "github.com/lib/pq"
	intNull := sql.NullInt16{}
	stringNull := sql.NullString{}
	boolNull := sql.NullBool{}

	db := GetConnection()
	defer db.Close()

	//Para traer informacion de la base de datos usamos Query
	rows, err := db.Query(q)

	if err != nil {
		log.Fatalf("Error en la consulta en la tabla de estudiantes => %s", err)
		return
	}
	//siempre se debe de cerrar todas las conexiones
	defer rows.Close()

	for rows.Next() {
		e := Estudiantes{}

		err = rows.Scan(
			&e.ID,
			&stringNull,
			&intNull,
			&boolNull,
			&e.CreatedAt,
			&timeNull,
		)
		if err != nil {
			log.Fatalf("Error al obtener los rows => %s", err.Error())
			return
		}

		// Asignamos los valores que se crearon a el slice de e

		e.UpdateAp = timeNull.Time
		e.Name = stringNull.String
		e.Age = int16(intNull.Int16)
		e.Active = boolNull.Bool

		estudiantes = append(estudiantes, e)
	}

	return estudiantes, nil
}

// funcion para la actualizacion de un estudiante
func Actualizar(e Estudiantes) error {
	q := `UPDATE estudiantes SET
	      name = $1, age = $2, active = $3, updated_at = now()
		  WHERE id = $4`

	db := GetConnection()
	defer db.Close()

	stm, err := db.Prepare(q)

	if err != nil {
		log.Fatalf("Error en la actualización de los datos => %s", err.Error())

	}
	defer stm.Close()

	//obtenemos la cantidad de filas que fueron afectadas
	r, err := stm.Exec(e.Name, e.Age, e.Active, e.ID)

	i, _ := r.RowsAffected()

	//Validamos el registro de una fila para poder mostrar el error
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila para actualizar")
	}
	return nil
}

func Borrar(id int) error {
	d := `DELETE FROM estudiantes WHERE id = $1`

	db := GetConnection()
	defer db.Close()

	stm, err := db.Prepare(d)

	if err != nil {
		log.Fatalf("Error en la eliminación => %s", err.Error())
	}

	r, err := stm.Exec(id)

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: se esperaba eliminar una fila de la base de datos")
	}

	return nil
}
