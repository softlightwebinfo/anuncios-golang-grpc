package structs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	host       string
	userName   string
	password   string
	name       string
	port       string
	Connection *sql.DB
}

func (then *Database) New(host string, userName string, password string, name string, port string) *Database {
	then.host = host
	then.userName = userName
	then.password = password
	then.name = name
	then.port = port
	return then
}

func (then *Database) Connect() *Database {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		then.host, then.port, then.userName, then.password, then.name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	then.Connection = db
	return then
}
