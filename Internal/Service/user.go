package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	dto "MVP/DTO"

	"github.com/google/uuid"
)

type UserService struct {
	users map[string]*dto.User
	file  string
	mutex sync.Mutex
}

// NewUserService 初始化 UserService
func NewUserService(file string) *UserService {
	us := &UserService{
		users: make(map[string]*dto.User),
		file:  file,
	}
	us.loadFromFile()
	return us
}

// 添加新用户（系统自动生成 address，不需要客户端提供）
func (us *UserService) AddUser(username string, balance float64) (*dto.User, error) {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	// 生成唯一 ID 和 address
	id := uuid.New().String()
	address := "ADDR_" + id[:8]

	user := &dto.User{
		ID:       id,
		Username: username,
		Address:  address,
		Balance:  balance,
	}

	us.users[id] = user
	if err := us.saveToFile(); err != nil {
		return nil, err
	}
	return user, nil
}

// 获取所有用户
func (us *UserService) GetAllUsers() []*dto.User {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	list := []*dto.User{}
	for _, u := range us.users {
		list = append(list, u)
	}
	return list
}

// 保存到 JSON
func (us *UserService) saveToFile() error {
	bytes, err := json.MarshalIndent(us.users, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(us.file, bytes, 0644)
}

// 从 JSON 加载
func (us *UserService) loadFromFile() error {
	if _, err := os.Stat(us.file); os.IsNotExist(err) {
		return nil
	}
	bytes, err := ioutil.ReadFile(us.file)
	if err != nil {
		return err
	}
	var data map[string]*dto.User
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}
	us.users = data
	return nil
}
