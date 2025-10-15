package handler

import (
	dto "MVP/DTO"
	service "MVP/Internal/Service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BlockchainHandler struct {
	bc  *service.BlockchainService
	pow *service.PowService
}

// 构造函数（注入区块链和PoW服务）
func NewBlockchainHandler(bc *service.BlockchainService, pow *service.PowService) *BlockchainHandler {
	return &BlockchainHandler{
		bc:  bc,
		pow: pow,
	}
}

/////////////////////////
// 区块和交易相关接口
/////////////////////////

// GET /api/v1/get_blocks
func (h *BlockchainHandler) GetBlocks(c *fiber.Ctx) error {
	blocks := h.bc.GetChain()
	return c.JSON(blocks)
}

// POST /api/v1/add_transaction
func (h *BlockchainHandler) AddTransaction(c *fiber.Ctx) error {
	var tx dto.Transaction
	if err := c.BodyParser(&tx); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "无效交易"})
	}
	if tx.From == "" || tx.To == "" || tx.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "交易字段不合法"})
	}

	newTx := h.bc.AddTransaction(&tx)
	return c.JSON(fiber.Map{"message": "交易已创建", "transaction": newTx})
}

// GET /api/v1/transactions
func (h *BlockchainHandler) GetPendingTransactions(c *fiber.Ctx) error {
	txs := h.bc.GetPendingTransactions()
	return c.JSON(txs)
}

// POST /api/v1/mine 验证待交易池的交易 + PoW 挖矿 + 验证 + 持久化
func (h *BlockchainHandler) MineTransactions(c *fiber.Ctx) error {
	pendingTx := h.bc.GetPendingTransactions()
	if len(pendingTx) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "没有可处理的交易"})
	}

	// 验证交易
	validTx := h.bc.ValidateTransactions(pendingTx)
	if len(validTx) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "不是有效的交易"})
	}

	// 创建新区块
	newBlock := h.bc.NewBlock(validTx)

	// 异步挖矿
	resultChan := h.pow.MineBlockAsync(newBlock)

	select {
	case minedBlock := <-resultChan:
		if !h.pow.ValidateBlock(minedBlock) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "工作量证明验证失败"})
		}

		// 添加到区块链
		h.bc.AddBlock(minedBlock)

		// 应用交易
		err := h.bc.ApplyTransactionsDirectly(minedBlock.Transactions)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "挖矿成功并应用交易",
			"block":   minedBlock,
		})

	case <-time.After(30 * time.Second):
		return c.JSON(fiber.Map{"message": "正在进行挖矿，请稍后再查看"})
	}
}
