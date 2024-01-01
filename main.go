package main

import (
	"task5-pbi-btpns-holidmuhamadsalman/database"
	"task5-pbi-btpns-holidmuhamadsalman/router"
)

func main() {
	database.ConnectDB()
	router := router.SetupRouter()

	router.Run(":8088")
}