package service

import (
	dto "MVP/DTO"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// PowService 提供挖矿和验证功能
type PowService struct{}

// NewPowService 构造函数
func NewPowService() *PowService {
	return &PowService{}
}

// MineBlock 对给定区块进行 PoW 挖矿
// 通过不断增加 Nonce 直到哈希前缀满足难度
func (p *PowService) MineBlock(block *dto.Block) {
	prefix := strings.Repeat("0", block.Difficulty)
	for {
		block.Timestamp = time.Now().Unix()
		hash := p.calculateHash(block)
		if strings.HasPrefix(hash, prefix) {
			block.Hash = hash
			break
		}
		block.Nonce++
	}
}

// 验证区块 PoW 是否有效
func (p *PowService) ValidateBlock(block *dto.Block) bool {
	prefix := strings.Repeat("0", block.Difficulty)
	hash := p.calculateHash(block)
	return strings.HasPrefix(hash, prefix) && hash == block.Hash
}

// 计算区块哈希
func (p *PowService) calculateHash(block *dto.Block) string {
	record := fmt.Sprintf("%d%d%s%d", block.Index, block.Timestamp, block.PrevHash, block.Nonce)
	for _, tx := range block.Transactions {
		record += fmt.Sprintf("%s%s%f", tx.From, tx.To, tx.Amount)
	}
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}
