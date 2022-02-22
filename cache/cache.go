package cache

type Cache struct {
	Prefix string
	Name string
}

//get cache key
func GetCacheKey(cache Cache) string {
	return cache.Prefix+":"+cache.Name
}
