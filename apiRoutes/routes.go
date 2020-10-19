package apiRoutes

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

type ResponseData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

var resp = []ResponseData{
	ResponseData{
		Id:   "1",
		Name: "test1",
		Age:  "22",
	},
	ResponseData{
		Id:   "2",
		Name: "test2",
		Age:  "23",
	},
}

func GetAllData(c *fiber.Ctx) error {

	log.Println(resp)
	return c.JSON(resp)

}
func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: "invalid id passed"})

	}
	if atoi >= len(resp) {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: "Index out of bounds|| not found"})

	}
	return c.JSON(resp[atoi])
}
