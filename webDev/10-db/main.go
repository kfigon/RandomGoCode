package main

import (
	"database/sql"
	"log"
	_ "github.com/jackc/pgx/stdlib"
)

type person struct {
	id int
	creationDate string
	name string
}

func main() {
	connectionString := "postgresql://localhost:5432/mydb?user=myuser&password=mypass"
	_, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal("error during opening: ", err)
	}
	log.Println("all gut")
}