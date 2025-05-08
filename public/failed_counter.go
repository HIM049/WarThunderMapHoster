package public

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type failedCounter struct {
	*cache.Cache
}

func InitFailedCounter() {
	FailedCounter.Cache = cache.New(30*time.Minute, 30*time.Minute)
}

func (fc *failedCounter) Get(ip string) int {
	count := 0
	value, found := fc.Cache.Get(ip)
	if found {
		count = value.(int)
	}
	return count
}

func (fc *failedCounter) Add(ip string) {
	var count int
	value, found := fc.Cache.Get(ip)
	if found {
		count = value.(int)
	}

	fc.Set(ip, count+1, cache.DefaultExpiration)
}

func (fc *failedCounter) Delete(ip string) {
	fc.Cache.Delete(ip)
}
