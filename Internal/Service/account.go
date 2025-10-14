package service

import (
	dto "MVP/DTO"
	"sync"
)

// AccountService 管理账户余额
type AccountService struct {
	balances map[string]float64
	mutex    sync.Mutex
}

// NewAccountService 构造函数
func NewAccountService() *AccountService {
	return &AccountService{
		balances: make(map[string]float64),
	}
}

// GetBalance 获取账户余额
func (s *AccountService) GetBalance(address string) float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if bal, ok := s.balances[address]; ok {
		return bal
	}
	return 0
}

// ApplyTransaction 应用交易，更新账户余额
func (s *AccountService) ApplyTransaction(tx *dto.Transaction) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 检查付款方余额
	if tx.From != "SYSTEM" && s.balances[tx.From] < tx.Amount {
		return false // 余额不足
	}

	// 扣款
	if tx.From != "SYSTEM" {
		s.balances[tx.From] -= tx.Amount
	}

	// 收款
	s.balances[tx.To] += tx.Amount

	return true
}

// 获取所有账户余额
func (s *AccountService) GetAllBalances() []*dto.Account {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var accounts []*dto.Account
	for addr, bal := range s.balances {
		accounts = append(accounts, &dto.Account{
			Address: addr,
			Balance: bal,
		})
	}
	return accounts
}
