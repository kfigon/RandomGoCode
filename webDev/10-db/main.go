package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://localhost:5432/mydb?user=myuser&password=mypass")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	readGreeting(conn)
}

func readGreeting(conn *pgx.Conn) {
	var greeting string
	err := conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
		return
	}

	log.Println(greeting)
}