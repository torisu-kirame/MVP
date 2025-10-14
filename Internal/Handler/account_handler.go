package handler

import (
	dto "MVP/DTO"
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

// POST /accounts/transaction 应用交易
func (h *AccountHandler) ApplyTransaction(c *fiber.Ctx) error {
	var tx dto.Transaction
	if err := c.BodyParser(&tx); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "无效交易"})
	}

	if tx.Amount <= 0 || tx.To == "" || tx.From == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "无效字段"})
	}

	ok := h.accountService.ApplyTransaction(&tx)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "余额不足"})
	}

	return c.JSON(fiber.Map{"message": "交易已应用", "transaction": tx})
}
