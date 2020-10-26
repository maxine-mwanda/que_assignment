package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func Connecttodb () (connection *sql.DB) {
	dburi :=os.Getenv("DBURI")
	connection, err:= sql.Open("mysql", dburi)
	if err !=nil {
		log.Println("Unable to connect to db", err)
		os.Exit(3)
	}

	log.Println("Succesfully connected")
	connection.SetMaxOpenConns(100)
	connection.SetMaxIdleConns(20)
	connection.SetConnMaxIdleTime(time.Second*10)
	return connection
}

func SaveToDb(db *sql.DB, counter int, timeStamp int64) {
	query := "INSERT INTO q (counter, time_stamp) VALUES (?, ?)"
	_, err:= db.Exec(query, counter, timeStamp)
	if err != nil {
		fmt.Println("unable to insert because", err)
		return
	}
	log.Printf("Saving to db. counter: %d, time_stamp: %d", counter, timeStamp)
}