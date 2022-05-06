package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

// Set a key/value
func StringSet(key string, data interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// SET if Not eXists
func StringSetNX(key string, data interface{}, seconds int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SETNX", key, value, seconds)
	if err != nil {
		return err
	}

	return nil
}

//SET key value EXPIRE key seconds
func StringSetEX(key string, data interface{}, seconds int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SETEX", key, value, seconds)
	if err != nil {
		return err
	}

	return nil
}

// Get get a key value
func StringGet(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

//incr key
func StringIncr(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("INCR", key)
	if err != nil {
		return err
	}

	return nil
}

//incr key by value
func StringIncrBy(key string, value int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("INCRBY", key, value)
	if err != nil {
		return err
	}

	return nil
}

//DECR key
func StringDecr(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("DECR", key)
	if err != nil {
		return err
	}

	return nil
}

//DECR key by value
func StringDecrBy(key string, value int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("DECRBY", key, value)
	if err != nil {
		return err
	}

	return nil
}
