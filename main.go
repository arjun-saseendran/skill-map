package main

import (
	"github.com/arjun-saseendran/skill-map/db"
	"github.com/arjun-saseendran/skill-map/handlers"
	"github.com/arjun-saseendran/skill-map/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDB()
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	userService := service.NewUserService()

	userHandler := handlers.NewUserHandlerFrom(userService)
	userHandler.RegisterEndpoints(router)

	router.Run()

}
