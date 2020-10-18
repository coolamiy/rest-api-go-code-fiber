package apiRoutes

import (
	"github.com/gofiber/fiber"
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
func GetAllData(c *fiber.Ctx) {

	log.Println(resp)
	_ = c.JSON(resp)
}
func GetOne(c *fiber.Ctx) {
	id := c.Params("id")
	atoi, err := strconv.Atoi(id)
	if err !=nil  {
		_ = c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: "invalid id passed"})
		return
	}
	if atoi >= len(resp) {
		_ = c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: "Index out of bounds|| not found"})
		return
	}
		_ = c.JSON(resp[atoi])

}
