package gredis

import (
	"github.com/gomodule/redigo/redis"
)

// Delete delete a kye
func KeyDel(key string) (bool, error) {
	conn := GetConn()
	return redis.Bool(conn.Do("DEL", key))
}

// Exists check a key
func KeyExists(key string) bool {
	conn := GetConn()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

//EXPIRE key seconds
func KeyExpire(key string, seconds int) bool {
	conn := GetConn()
	exists, err := redis.Bool(conn.Do("EXPIRE", key, seconds))
	if err != nil {
		return false
	}

	return exists
}

//EXPIRE key unix timestamp
func KeyExpireAt(key string, timeStamp int) bool {
	conn := GetConn()
	exists, err := redis.Bool(conn.Do("EXPIREAT", key, timeStamp))
	if err != nil {
		return false
	}

	return exists
}

func KeyAllKeys(key string) ([]byte, error) {
	conn := GetConn()

	reply, err := redis.Bytes(conn.Do("KEYS", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}
