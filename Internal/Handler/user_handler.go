package handler

import (
	service "MVP/Internal/Service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// POST /users 添加用户
func (h *UserHandler) AddUser(c *fiber.Ctx) error {
	var req struct {
		Username string  `json:"username"`
		Balance  float64 `json:"balance"` // 可选，默认 0
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username is required"})
	}

	user, err := h.userService.AddUser(req.Username, req.Balance)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// 返回完整用户信息，包括系统自动生成的 address
	return c.JSON(user)
}

// GET /users 获取所有用户
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users := h.userService.GetAllUsers()
	return c.JSON(users)
}
