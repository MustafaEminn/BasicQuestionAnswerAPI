package main

import (
	"forLearnCurrent/app"
	"forLearnCurrent/configs"
	"forLearnCurrent/repository"
	"forLearnCurrent/services"
	"github.com/gofiber/fiber/v2"
)

func main(){
	appRoute := fiber.New()
	configs.ConnectDB()

	dbClient := configs.GetCollection(configs.DB, "mains")
	MainRepositoryDB := repository.NewMainRepositoryDB(dbClient)

	main := app.MainHandler{Service: services.NewMainService(MainRepositoryDB)}

	appRoute.Post("/api/main", main.CreateMain)
	appRoute.Get("/api/main/:id", main.GetByIdMain)
	appRoute.Get("/api/mains", main.GetAllMain)
	appRoute.Delete("/api/todo/:id", main.DeleteMain)

	appRoute.Listen(":8080")
}
