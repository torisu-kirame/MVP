package handler

import (
	dto "MVP/DTO"
	service "MVP/Internal/Service"

	"github.com/gofiber/fiber/v2"
)

// BlockchainHandler 统一管理区块链和交易相关接口
type BlockchainHandler struct {
	bc *service.BlockchainService
}

// 构造函数
func NewBlockchainHandler(bc *service.BlockchainService) *BlockchainHandler {
	return &BlockchainHandler{bc: bc}
}

/////////////////////////
// 区块相关接口
/////////////////////////

// GET /blocks 获取区块链
func (h *BlockchainHandler) GetBlocks(c *fiber.Ctx) error {
	blocks := h.bc.GetChain()
	resBlocks := make([]*dto.Block, len(blocks))
	for i, b := range blocks {
		resBlocks[i] = &dto.Block{
			Index:        b.Index,
			Timestamp:    b.Timestamp,
			Transactions: b.Transactions,
			PrevHash:     b.PrevHash,
			Hash:         b.Hash,
			Nonce:        b.Nonce,
			Difficulty:   b.Difficulty,
		}
	}

	resp := dto.BlockchainResponse{
		Chain:      resBlocks,
		State:      nil,             // 可扩展系统状态
		Difficulty: h.bc.Difficulty, // 当前链默认难度
		PendingTx:  h.bc.PendingTx,  // 当前交易池
	}

	return c.JSON(resp)
}

// POST /blocks 添加新区块（打包交易生成新区块）
func (h *BlockchainHandler) AddBlock(c *fiber.Ctx) error {
	var req dto.BlockchainResponse // 或自定义 AddBlockRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if len(req.PendingTx) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Transactions are required"})
	}

	// 添加交易到交易池
	for _, tx := range req.PendingTx {
		h.bc.AddTransaction(tx)
	}

	// 打包交易生成新区块
	block := h.bc.NewBlock(h.bc.PendingTx)

	res := &dto.Block{
		Index:        block.Index,
		Timestamp:    block.Timestamp,
		Transactions: block.Transactions,
		PrevHash:     block.PrevHash,
		Hash:         block.Hash,
		Nonce:        block.Nonce,
		Difficulty:   block.Difficulty,
	}

	return c.JSON(res)
}

/////////////////////////
// 交易相关接口
/////////////////////////

// GET /transactions 获取交易池
func (h *BlockchainHandler) GetPendingTransactions(c *fiber.Ctx) error {
	txs := h.bc.GetPendingTransactions()
	return c.JSON(txs)
}
