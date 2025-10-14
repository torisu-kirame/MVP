package service

import (
	dto "MVP/DTO"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type BlockchainService struct {
	Chain      []*dto.Block
	PendingTx  []*dto.Transaction
	mutex      sync.Mutex
	file       string
	Difficulty int
}

// 创建服务/创世区块
func NewBlockchainService(file string, difficulty int) *BlockchainService {
	bc := &BlockchainService{
		file:       file,
		Difficulty: difficulty,
		PendingTx:  []*dto.Transaction{},
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		genesis := &dto.Block{
			Index:        0,
			Timestamp:    time.Now().Unix(),
			Transactions: []*dto.Transaction{},
			PrevHash:     "0",
			Difficulty:   difficulty,
			Nonce:        0,
		}
		genesis.Hash = calculateHash(genesis)
		bc.Chain = []*dto.Block{genesis}
		bc.saveToFile()
	} else {
		bc.loadFromFile()
	}

	return bc
}

// 计算区块哈希
func calculateHash(block *dto.Block) string {
	record := fmt.Sprint(block.Index) + fmt.Sprint(block.Timestamp) + block.PrevHash + fmt.Sprint(block.Nonce)
	for _, tx := range block.Transactions {
		record += tx.From + tx.To + fmt.Sprint(int(tx.Amount*100))
	}
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// 创建新区块
func (bc *BlockchainService) NewBlock(transactions []*dto.Transaction) *dto.Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	if len(bc.PendingTx) == 0 {
		return nil // 没有交易就不生成块
	}

	prevBlock := bc.Chain[len(bc.Chain)-1]
	block := &dto.Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Difficulty:   bc.Difficulty,
		Nonce:        0,
	}
	block.Hash = calculateHash(block)
	bc.Chain = append(bc.Chain, block)
	bc.saveToFile()

	// 清空待处理交易池
	bc.PendingTx = []*dto.Transaction{}
	// 持久化
	bc.saveToFile()

	return block
}

func (bc *BlockchainService) AddBlock(block *dto.Block) {
	bc.Chain = append(bc.Chain, block)
	bc.PendingTx = []*dto.Transaction{} // 清空交易池
	bc.saveToFile()                     // 写入 blockchain.json
}

// 获取区块链
func (bc *BlockchainService) GetChain() []*dto.Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	return bc.Chain
}

// 保存到文件
func (bc *BlockchainService) saveToFile() error {
	bytes, err := json.MarshalIndent(bc.Chain, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(bc.file, bytes, 0644)
}

// 从文件加载
func (bc *BlockchainService) loadFromFile() error {
	bytes, err := ioutil.ReadFile(bc.file)
	if err != nil {
		return err
	}
	var blocks []*dto.Block
	if err := json.Unmarshal(bytes, &blocks); err != nil {
		return err
	}
	bc.Chain = blocks
	return nil
}
