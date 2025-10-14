package handler

import (
	dto "MVP/DTO"
	service "MVP/Internal/Service"

	"github.com/gofiber/fiber/v2"
)

type PowHandler struct {
	powService *service.PowService
}

// 构造函数
func NewPowHandler(powService *service.PowService) *PowHandler {
	return &PowHandler{powService: powService}
}

// POST /pow/mine 对给定区块进行挖矿
func (h *PowHandler) MineBlock(c *fiber.Ctx) error {
	var block dto.Block
	if err := c.BodyParser(&block); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid block data"})
	}
	// 执行挖矿
	h.powService.MineBlock(&block)
	// 挖矿完成后立即验证PoW是否有效
	valid := h.powService.ValidateBlock(&block)

	response := fiber.Map{
		"message": "Block mining completed",
		"block":   block,
		"validation": fiber.Map{
			"valid":      valid,
			"hash":       block.Hash,
			"nonce":      block.Nonce,
			"difficulty": block.Difficulty,
		},
	}

	if !valid {
		response["warning"] = "Block was mined but PoW validation failed"
	}
	return c.JSON(response)
}
