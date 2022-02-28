package emitter

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

const (
	insertQuery = "INSERT INTO RequestLog (merchantId, request, response, addedOn) VALUES(?, ?, ?, ?)"
)

type Database struct {
	Db sql.DB
}

func InitDB(username string, password string, database string) Database {
	connStr := username + ":" + password + "@/" + database
	db, err := sql.Open("mysql", connStr)
	log.Info("Setting connection with database")
	if err != nil {
		fmt.Println("Cannot connect to database")
	}
	db.SetConnMaxLifetime(time.Minute * 30)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Error("Connection failed ", err)
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	log.Info("Connection with database established")
	return Database{Db: *db}
}

func (_this Database) Emit(event Event) error {
	log.Info("sending data to mysql")
	//TODO understand this mystery - if i give incorrect insert query
	err := _this.Db.Ping()
	if err != nil {
		log.Error("Connection failed ", err)
		return err
	}
	stmtIns, err := _this.Db.Prepare(insertQuery)
	if err != nil {
		log.Error("failed in preparing query ")
		return err
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(event.MID, event.Request, event.Response, "2022-02-20 21:45")
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println("failure in getting result, ", err)
		return err
	}
	fmt.Println("result is: ", n)
	return nil
}
