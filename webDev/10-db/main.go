package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
)

type person struct {
	id int
	creationDate string
	name string
}
func (p person) String() string {
	return fmt.Sprintf("id: %v, name: %v, date: %v", p.id, p.name, p.creationDate)
}

func main() {
	connectionString := "postgresql://myuser:mypass@localhost:5432/mydb"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal("error during opening: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("error during ping: ", err)
	}

	rows, err := db.Query("select id, creation_date, name from person")
	if err != nil {
		log.Fatal("Error during querying: ", err)
	}
	defer rows.Close()
	
	people := make([]person,0)
	for rows.Next() {
		p := person{}
		check(rows.Scan(&p.id, &p.creationDate, &p.name))

		people = append(people, p)
	}
	log.Println("all gut, got")
	for _, v := range people {
		log.Println(v)
	}
}

func check(err error) {
	if err != nil {
		log.Println("Got error: ", err)
	}
}