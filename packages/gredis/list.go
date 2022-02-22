package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

// Set a key/value
func ListLPush(key string, data interface{}) error {
	conn := GetConn()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = redis.Int(conn.Do("LPUSH", key, value))
	if err != nil {
		return err
	}

	return nil
}