package main

import (
	"fmt"

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
	skillService := service.NewSkillService()

	userHandler := handlers.NewUserHandlerFrom(userService)
	userHandler.RegisterEndpoints(router)

	skillHandler := handlers.NewSkillHandleFrom(skillService)
	skillHandler.RegisterEndpoints(router)

	err := router.Run()
	if err != nil {
		fmt.Println("sever failed to run.")
		return
	}

}
