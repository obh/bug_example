package main

import (
	"fmt"
	"os"

	_db "github.com/obh/bug_example/emitter"
	_route "github.com/obh/bug_example/route"

	"github.com/labstack/echo/v4"
)

func main() {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB")
	db := _db.InitDB(user, password, dbName)
	fmt.Println(db)

	e := echo.New()

	eventRoute := e.Group("v1")
	_route.ConfigureEventHandlerHTTP(eventRoute, &db)

	//go e.Start(cfg.Port)
	e.Logger.Fatal(e.Start(":7000"))
	//graceFullShutdown(e)
}
