package main

import (
	api "MVP/API"
	handler "MVP/Internal/Handler"
	service "MVP/Internal/Service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	blockchain_path := "data/blockchain.json"
	users_data_path := "data/users.json"

	app := fiber.New()

	// 初始化区块链服务
	const difficulty = 1 // 设置难度
	bcs := service.NewBlockchainService(blockchain_path, difficulty)
	pows := service.NewPowService()
	accounts := service.NewAccountService(users_data_path)
	userService := service.NewUserService(users_data_path)

	// 初始化路由
	bcHandler := handler.NewBlockchainHandler(bcs, pows)
	accountHandler := handler.NewAccountHandler(accounts)
	userHandler := handler.NewUserHandler(userService)

	// 注册路由
	api.BlockHandlerRoutes(app, bcHandler)
	api.AccountRoutes(app, accountHandler)
	api.UserRoutes(app, userHandler)

	// 启动服务
	log.Println("区块链服务运行于 :http://localhost:8080/")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("服务器故障: %v", err)
	}

}
