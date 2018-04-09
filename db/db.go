/*
Package db implemetns logic to write logs in local postgres database
*/

package db

import (
	"database/sql"
	"fmt"
	"log"
	"tbot/config"

	_ "github.com/lib/pq"
)

type Logs struct {
	Id    int
	Login string
	Text  string
	Chat  string
	Time  string
}

var db *sql.DB

// Init open connection to local PG database
func init() {
	var query string
	var err error
	query = fmt.Sprintf("postgres://%s:%s@localhost/%s", config.DB_USER, config.DB_PASS, config.DB_NAME)
	db, err = sql.Open("postgres", query)

	if err != nil {
		log.Fatal(err)
	}
	// check database connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

//logToString prepeare string with log content to added it in database
func logToString(l *Logs) string {
	var s string
	s = fmt.Sprintf("'%s','%s','%s','%s'", l.Login, l.Text, l.Chat, l.Time)
	log.Println(s)
	return s
}

// AddMessageToLog write logs in local postgres database

func AddMessageToLog(l *Logs) {
	var query string
	query = fmt.Sprintf("INSERT INTO logs VALUES (%s)", logToString(l))
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
