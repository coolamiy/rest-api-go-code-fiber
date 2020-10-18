package apiRoutes

import (
	"github.com/gofiber/fiber"
	"log"
)

func GetAllData(c *fiber.Ctx){
	log.Print("this is get all data")
	c.SendString("Get all data Route")
}
func  GetOne(c *fiber.Ctx)  {
	c.SendString("send one ")

}