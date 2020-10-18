package main

import (
	"github.com/coolamiy/go-rest-api/apiRoutes"
	"github.com/gofiber/fiber"
	"log"
)

func handlerHome(c *fiber.Ctx) {

	c.SendString("pong")

}
func setupRoutes(app *fiber.App)  {
	app.Get("/", handlerHome)
	app.Get("/data",apiRoutes.GetAllData)
	app.Get("/data/:id", apiRoutes.GetOne)
	
}

func main() {
	log.Println("Starting App..")
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
