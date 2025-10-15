package service

import (
	dto "MVP/DTO"
	"encoding/json"
	"fmt"
	"os"
)

// 添加交易到交易池
func (bc *BlockchainService) AddTransaction(tx *dto.Transaction) *dto.Transaction {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	bc.PendingTx = append(bc.PendingTx, tx)
	return tx
}

// 获取当前交易池
func (bc *BlockchainService) GetPendingTransactions() []*dto.Transaction {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	return bc.PendingTx
}

// 验证交易
func (bc *BlockchainService) ValidateTransactions(txs []*dto.Transaction) []*dto.Transaction {
	validTx := []*dto.Transaction{}
	for _, tx := range txs {
		if tx.Amount > 0 && tx.From != "" && tx.To != "" {
			// 可扩展账户余额检查
			validTx = append(validTx, tx)
		}
	}
	return validTx
}

// 遍历交易，更新 user.json
func (h *BlockchainService) ApplyTransactionsDirectly(transactions []*dto.Transaction) error {
	// 读取用户数据
	bytes, err := os.ReadFile("data/users.json")
	if err != nil {
		return err
	}

	var users map[string]*dto.User
	if err := json.Unmarshal(bytes, &users); err != nil {
		return err
	}

	// 辅助函数：根据 Address 查找用户
	findUserByAddress := func(address string) (*dto.User, error) {
		for _, u := range users {
			if u.Address == address {
				return u, nil
			}
		}
		return nil, fmt.Errorf("找不到用户: %s", address)
	}

	// 遍历交易应用到用户余额
	for _, tx := range transactions {
		// 系统奖励交易，不扣款
		if tx.From != "SYSTEM" {
			sender, err := findUserByAddress(tx.From)
			if err != nil {
				return err
			}
			if sender.Balance < tx.Amount {
				return fmt.Errorf("%s 余额不足", tx.From)
			}
			sender.Balance -= tx.Amount
		}

		receiver, err := findUserByAddress(tx.To)
		if err != nil {
			return err
		}
		receiver.Balance += tx.Amount
	}

	// 保存回 user.json
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("data/users.json", data, 0644)
}
