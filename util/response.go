package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Handler struct {
	response *Response
	c        *fiber.Ctx
}

func CreateResponse(c *fiber.Ctx) *Handler {
	return &Handler{
		response: &Response{},
		c:        c,
	}
}

func (f *Handler) SendResponse(statusCode int, message string, data interface{}) error {

	if statusCode > 0 {
		f.c.Status(statusCode)
	}
	f.response.Data = data
	f.response.Message = message

	return f.c.JSON(f.response)
}
