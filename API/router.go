package api

import (
	handler "MVP/Internal/Handler"

	"github.com/gofiber/fiber/v2"
)

func BlockHandlerRoutes(app *fiber.App, blockHandler *handler.BlockchainHandler) {
	v1 := app.Group("/api/v1")

	// 区块链相关路由
	v1.Get("/get_blocks", blockHandler.GetBlocks)                    // 获取区块链
	v1.Post("/add_transaction", blockHandler.AddTransaction)         // 生成交易
	v1.Post("/mine", blockHandler.MineTransactions)                  //挖矿
	v1.Get("/get_transactions", blockHandler.GetPendingTransactions) // 获取交易池
}

func AccountRoutes(app *fiber.App, handler *handler.AccountHandler) {
	v1 := app.Group("/api/v1/accounts")

	v1.Get("/", handler.GetAllBalances)     // 查询所有账户余额
	v1.Get("/:address", handler.GetBalance) // 查询单个账户余额
}

func UserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	v1 := app.Group("/api/v1/users")
	v1.Post("/add", userHandler.AddUser)     // 添加用户
	v1.Get("/get_all", userHandler.GetUsers) // 获取所有用户 a
}
