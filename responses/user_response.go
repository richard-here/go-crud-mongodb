package responses

import "github.com/gofiber/fiber/v2"

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func ErrorUserResponse(status int, err error) UserResponse {
	return UserResponse{
		Status:  status,
		Message: "error",
		Data:    &fiber.Map{"data": err.Error()},
	}
}
