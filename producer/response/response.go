package response

import "github.com/gofiber/fiber/v2"

type (
	MsgResponse struct {
		Message string `json:"message"`
	}
)

func ErrResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(&MsgResponse{message})
}

func SuccessResponse(c *fiber.Ctx, statusCode int, data any) error {
	return c.Status(statusCode).JSON(data)

}
