package main

import (
	"fmt"

	"github.com/jackc/pgx"
)

func main() {
	conn, err := NewConn()
	if err != nil {
		fmt.Println("Unable to connect to database", err)
		return
	}
	defer conn.Close()

	b, err := querySingle(conn, func(r *pgx.Row, b *blog) error {
		return r.Scan(&b.id, &b.title, &b.text)
	},"select id, title, text from blog limit 1")
	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)

	
	blogs, err := queryMany(conn, func(r *pgx.Rows, b *blog) error {
		return r.Scan(&b.id, &b.title, &b.text)
	},"select id, title, text from blog")
	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(blogs)
}

func NewConn() (*pgx.Conn, error) {
	return pgx.Connect(pgx.ConnConfig{
		Host: "localhost",
		User: "postgres",
		Password: "postgres",
		Port: 5432,
	})
}

type mappingSingle[T any] func(*pgx.Row, *T) error
type mappingMany[T any] func(*pgx.Rows, *T) error

type blog struct {
	id int
	title string
	text string
}

func querySingle[T any](conn *pgx.Conn, m mappingSingle[T], query string, args ...interface{}) (T, error) {
	var out T
	row := conn.QueryRow(query, args...)
	err := m(row, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func queryMany[T any](conn *pgx.Conn, m mappingMany[T], query string, args ...interface{}) ([]T, error) {
	var out []T
	rows, err := conn.Query(query, args...)
	if err != nil {
		return out, err
	}

	for rows.Next() {
	  var v T
	  err = m(rows, &v)
	  if err != nil {
		  return out, err
	  }
	  out = append(out, v)
	}

	return out, nil
}