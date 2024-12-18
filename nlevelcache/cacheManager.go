package main

type CacheManager struct {
	CacheLevels []ICache
}

func NewCacheManager(n int) CacheManager {
	CacheLevels := make([]ICache, n)
	for i := 0; i < n; i++ {
		CacheLevels[i] = NewCache(i)
		if i > 0 {
			CacheLevels[i-1].Next(CacheLevels[i])
		}
	}

	return CacheManager{CacheLevels: CacheLevels}
}

func (cm *CacheManager) Read(key string) (*string, error) {
	return cm.CacheLevels[0].Read(key)
}
