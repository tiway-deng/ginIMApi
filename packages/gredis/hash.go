package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
)

// Set a  HASH key/value
func HashSet(key string, field string, data interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("HSET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// Get get a HASH key value
func HashGet(key string, field string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("HGET", key, field))
	log.Print(reply, err)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a HASH kye
func HashKeyDel(key string) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Int(conn.Do("HDEL", key))
}

//incr value by hash key
func HashIncrBy(key string, field string, value int) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Int(conn.Do("HINCRBY", key, field, value))
}
