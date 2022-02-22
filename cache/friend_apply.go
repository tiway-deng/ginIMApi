package cache

import (
	"ginIMApi/packages/gredis"
)

func NewApplyFriend() *Cache {
	return &Cache{
		Prefix: "im",
		Name:   "apply-friend",
	}
}

func (c Cache) IncrApplyFriendUnRead(friendId string) (int, error) {
	return gredis.HashIncrBy(GetCacheKey(c), friendId, 1)
}

func (c Cache) GetApplyFriendUnRead(friendId string) ([]byte, error) {
	return gredis.HashGet(GetCacheKey(c), friendId)
}
