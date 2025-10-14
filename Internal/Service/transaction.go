package service

import (
	dto "MVP/DTO"
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
