package dto

type Block struct {
	Index        int64          `json:"index"`
	Timestamp    int64          `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
	PrevHash     string         `json:"prev_hash"`
	Hash         string         `json:"hash"`
	Nonce        int64          `json:"nonce"`
	Difficulty   int            `json:"difficulty"`
}

type BlockchainResponse struct {
	Chain      []*Block       `json:"chain"`
	State      interface{}    `json:"state"` // 这里简化为 interface{}，可自定义
	Difficulty int            `json:"difficulty"`
	PendingTx  []*Transaction `json:"pending_tx"`
}

// Transaction 交易结构
type Transaction struct {
	TxID   string  `json:"txid"`   // 交易 ID（可用 UUID 或 Hash）
	From   string  `json:"from"`   // 付款方账户地址
	To     string  `json:"to"`     // 收款方账户地址
	Amount float64 `json:"amount"` // 转账金额
}

// UTXO 模型中的账户余额
type Account struct {
	Address string  `json:"address"` // 账户地址
	Balance float64 `json:"balance"` // 账户余额
}
