package main

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/routers"
	workers "localArtisans/worker"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
	configs.ConnectToRedis()
}

func main() {
	db := configs.GetDB()
	db.AutoMigrate(&database.User{}, &database.Artisans{}, &database.Categories{}, &database.Products{}, &database.Promos{}, &database.Reviews{}, &database.Carts{}, &database.CartInformations{}, &database.Orders{}, &database.OrderItems{}, &database.LogActivity{})
	
	go workers.StartLogWorker()

	r := routers.RoutersConfiguration()
	r.Run(":3000")
}