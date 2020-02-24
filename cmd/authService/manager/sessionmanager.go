package manager

import (
	"github.com/go-redis/redis/v7"
)

type SessionManagerInterface interface {
	GetUserIDByToken(token string) (int64, error)
	SetTokenForUserID(token string, userID int64) error
	DeleteToken(token string) error
}

type RedisManager struct {
	client *redis.Client
}

func NewRedisManager(addr, password string, db int) *RedisManager {
	client := redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db})
	return &RedisManager{client: client}
}

func (r RedisManager) GetUserIDByToken(token string) (int64, error) {
	userID, err := r.client.Get(token).Int64()
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r RedisManager) SetTokenForUserID(token string, userID int64) error {
	err := r.client.Set(token, userID, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisManager) DeleteToken(token string) error {
	err := r.client.Del(token).Err()
	if err != nil {
		return err
	}
	return nil
}
