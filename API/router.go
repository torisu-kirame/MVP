package api

import (
	handler "MVP/Internal/Handler"

	"github.com/gofiber/fiber/v2"
)

func BlockHandlerRoutes(app *fiber.App, blockHandler *handler.BlockchainHandler) {
	v1 := app.Group("/api/v1")

	// 区块链相关路由
	v1.Get("/get_blocks", blockHandler.GetBlocks)                    // 获取区块链
	v1.Post("/post_blocks", blockHandler.AddBlock)                   // 生成交易
	v1.Get("/get_transactions", blockHandler.GetPendingTransactions) // 获取交易池
}

func PowRoutes(app *fiber.App, powHandler *handler.PowHandler) {
	v1 := app.Group("/api/v1/pow")
	v1.Post("/mine", powHandler.MineBlock) // 挖矿
}

func AccountRoutes(app *fiber.App, handler *handler.AccountHandler) {
	v1 := app.Group("/api/v1/accounts")

	v1.Get("/", handler.GetAllBalances)               // 查询所有账户余额
	v1.Get("/:address", handler.GetBalance)           // 查询单个账户余额
	v1.Post("/transaction", handler.ApplyTransaction) // 应用交易
}
