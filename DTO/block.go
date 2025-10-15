package dto

type Block struct {
	Index        int64          `json:"index"`        //区块高度
	Timestamp    int64          `json:"timestamp"`    //时间戳
	Transactions []*Transaction `json:"transactions"` //交易列表
	PrevHash     string         `json:"prev_hash"`    //前区块哈希值
	Hash         string         `json:"hash"`         //当前哈希
	Nonce        int64          `json:"nonce"`        //随机数
	Difficulty   int            `json:"difficulty"`   //挖矿难度
}

// Transaction 交易结构
type Transaction struct {
	TxID   string  `json:"txid"`   // 交易 ID
	From   string  `json:"from"`   // 付款方账户地址
	To     string  `json:"to"`     // 收款方账户地址
	Amount float64 `json:"amount"` // 转账金额
}

// type BlockchainResponse struct {
// 	Chain []*Block    `json:"chain"` //区块链
// 	State interface{} `json:"state"` //当前链的全局状态
// 	// Difficulty int            `json:"difficulty"` //挖矿难度
// 	PendingTx []*Transaction `json:"pending_tx"` //未打包的交易池
// }
