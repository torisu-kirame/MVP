package handler

import (
	service "MVP/Internal/Service"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	accountService *service.AccountService
}

// 构造函数
func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

// GET /accounts/:address 查询账户余额
func (h *AccountHandler) GetBalance(c *fiber.Ctx) error {
	addr := c.Params("address")
	balance := h.accountService.GetBalance(addr)
	return c.JSON(fiber.Map{"address": addr, "balance": balance})
}

// GET /accounts 查询所有账户余额
func (h *AccountHandler) GetAllBalances(c *fiber.Ctx) error {
	accounts := h.accountService.GetAllBalances()
	return c.JSON(accounts)
}
