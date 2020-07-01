package main

import (
	"mysqlApp/api/master"
	"mysqlApp/config"
)

func main() {
	db := config.ConnectionDB()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
