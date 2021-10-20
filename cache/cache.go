package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

var MyCache CacheItf

type Getter func() interface{}

type CacheItf interface {
	Set(key string, data interface{}, expiration time.Duration) error
	Get(key string) ([]byte, error)
}

type AppCache struct {
	Client *cache.Cache
}

func (r *AppCache) Set(key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r.Client.Set(key, b, expiration)
	return nil
}

func (r *AppCache) Get(key string) ([]byte, error) {
	res, exist := r.Client.Get(key)
	if !exist {
		return nil, nil
	}

	resByte, ok := res.([]byte)
	if !ok {
		return nil, errors.New("Format is not arr of bytes")
	}

	return resByte, nil
}

func InitCache() {
	MyCache = &AppCache{
		Client: cache.New(5*time.Minute, 10*time.Minute),
	}
}
