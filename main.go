package main

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/routers"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
}

// @title Local Artisans API
// @version 1.0
// @description This is the API for Local Artisans
// @BasePath /api/v1
// @AuthPath /api/v1/auth -> coming soon
func main() {
	db := configs.GetDB()
	db.AutoMigrate(&database.User{}, &database.Artisans{}, &database.Categories{}, &database.Products{}, &database.Promos{}, &database.Reviews{}, &database.Carts{}, &database.CartInformations{})

	r := routers.RoutersConfiguration()
	r.Run(":3000")
}