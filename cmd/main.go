package main

import (
	api "MVP/API"
	handler "MVP/Internal/Handler"
	service "MVP/Internal/Service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// 初始化区块链服务
	const difficulty = 3
	bc := service.NewBlockchainService("data/blockchain.json", difficulty)
	powService := service.NewPowService()
	accountService := service.NewAccountService()

	// 初始化路由
	bcHandler := handler.NewBlockchainHandler(bc)
	powHandler := handler.NewPowHandler(powService)
	accountHandler := handler.NewAccountHandler(accountService)

	// 注册路由
	api.BlockHandlerRoutes(app, bcHandler)
	api.PowRoutes(app, powHandler)
	api.AccountRoutes(app, accountHandler)

	// 启动服务
	log.Println("Blockchain server running on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
