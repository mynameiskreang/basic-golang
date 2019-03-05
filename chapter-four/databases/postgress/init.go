package postgress

import (
	"github.com/jackc/pgx"
	"log"
)

var pgxConn *pgx.Conn

func init() {
	config := pgx.ConnConfig{Host: "172.17.0.2", Port: 5432, Database: "dvdrental", User: "postgres", Password: "toor123"}
	conn, err := pgx.Connect(config)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	pgxConn = conn
}
