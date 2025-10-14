package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	dto "MVP/DTO"
)

type AccountService struct {
	users map[string]*dto.User
	file  string
	mutex sync.Mutex
}

// 初始化
func NewAccountService(file string) *AccountService {
	s := &AccountService{
		users: make(map[string]*dto.User),
		file:  file,
	}
	s.loadFromFile()
	return s
}

// 获取余额
func (s *AccountService) GetBalance(address string) float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	bytes, err := os.ReadFile(s.file)
	if err != nil {
		return 0
	}

	var users map[string]*dto.User
	if err := json.Unmarshal(bytes, &users); err != nil {
		return 0
	}

	for _, u := range users {
		if u.Address == address {
			return u.Balance
		}
	}

	return 0
}

// 获取所有账户
func (s *AccountService) GetAllBalances() []*dto.User {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 从文件读取
	bytes, err := os.ReadFile(s.file)
	if err != nil {
		return []*dto.User{} // 出错返回空
	}

	var users map[string]*dto.User
	if err := json.Unmarshal(bytes, &users); err != nil {
		return []*dto.User{} // 出错返回空
	}

	list := []*dto.User{}
	for _, u := range users {
		list = append(list, u)
	}

	return list
}

// 从 JSON 加载
func (s *AccountService) loadFromFile() error {
	if _, err := os.Stat(s.file); os.IsNotExist(err) {
		return nil
	}
	bytes, err := ioutil.ReadFile(s.file)
	if err != nil {
		return err
	}
	var data map[string]*dto.User
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	s.users = data
	return nil
}
