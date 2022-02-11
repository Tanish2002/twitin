package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "weeb"
	password = "kt20020801"
	dbname   = "twitin"
)

var db *sql.DB

func init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//fmt.Println(psqlconn)
	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
	}
	db = conn
	// defer conn.Close()
}
