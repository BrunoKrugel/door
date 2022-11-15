package main

import (
	"github.com/BrunoKrugel/door/client"
	"github.com/BrunoKrugel/door/db"
	"github.com/labstack/echo"
)

func main() {
	// Start in memory database
	db.Init()

	// Start http server
	e := echo.New()
	e.GET("/door", client.PickDoor)
	e.Logger.Fatal(e.Start(":4554"))

}
