package main

import (
	"github.com/jackc/pgx"
)

type repo struct {
	conn *pgx.Conn
}

type blog struct {
	id int
	title string
	text string
}

func (r *repo) queryBlog(id int) (blog, error) {
	return	querySingle(r.conn, func(r *pgx.Row, b *blog) error {
		return r.Scan(&b.id, &b.title, &b.text)
	},
	"select id, title, text from blog where id=$1", id)
}

func (r *repo) queryAllBlogs() ([]blog, error) {
	return	queryMany(r.conn, func(r *pgx.Rows, b *blog) error {
		return r.Scan(&b.id, &b.title, &b.text)
	},
	"select id, title, text from blog")
}

func newRepo() (*repo, error) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host: "localhost",
		User: "postgres",
		Password: "postgres",
		Port: 5432,
	})

	if err != nil {
		return nil, err
	}
	return &repo{conn}, nil
}

type mappingSingle[T any] func(*pgx.Row, *T) error
type mappingMany[T any] func(*pgx.Rows, *T) error

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