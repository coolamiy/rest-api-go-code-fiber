package main

import (
	"github.com/coolamiy/go-rest-api/apiRoutes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/helmet/v2"
	"log"
)

func handlerHome(c *fiber.Ctx) error {

	c.SendString("pong")
	return nil

}
func SetupRoutes(app *fiber.App) {
	app.Get("/", handlerHome)
	app.Get("/data", apiRoutes.GetAllData)
	app.Get("/data/:id", apiRoutes.GetOne)

}

func main() {
	log.Println("Starting App..")
	app := fiber.New()
	//app := fiber.New(fiber.Config{
	//	Prefork:                   false,
	//	ServerHeader:              "",
	//	StrictRouting:             false,
	//	CaseSensitive:             false,
	//	Immutable:                 false,
	//	UnescapePath:              false,
	//	ETag:                      false,
	//	BodyLimit:                 0,
	//	Concurrency:               0,
	//	Views:                     nil,
	//	ReadTimeout:               0,
	//	WriteTimeout:              0,
	//	IdleTimeout:               0,
	//	ReadBufferSize:            0,
	//	WriteBufferSize:           0,
	//	CompressedFileSuffix:      "",
	//	ProxyHeader:               "",
	//	GETOnly:                   false,
	//	ErrorHandler:              nil,
	//	DisableKeepalive:          false,
	//	DisableDefaultDate:        false,
	//	DisableDefaultContentType: false,
	//	DisableHeaderNormalizing:  false,
	//	DisableStartupMessage:     true,
	//	ReduceMemoryUsage:         true,
	//})
	app.Use(helmet.New())
	app.Get("/dashboard", monitor.New())
	app.Use(logger.New())
	SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
