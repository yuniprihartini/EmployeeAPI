package main

import (
	"test/configs"
	"test/master"
)

func main() {
	db, _ := configs.ConnectionDB()
	router := configs.CreateRouter()
	master.Init(router, db)
	configs.RunServer(router)
}
