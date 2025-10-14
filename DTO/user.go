package dto

// User 用户信息 & UTXO 模型中的账户余额
type User struct {
	ID       string  `json:"id"`       // 用户唯一ID
	Username string  `json:"username"` // 用户名
	Address  string  `json:"address"`  // 账户地址
	Balance  float64 `json:"balance"`  // 账户余额
}
