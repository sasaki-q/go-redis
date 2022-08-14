package item

import (
	"time"

	"example.com/redis/config/cache"
)

type ItemService struct{}

func (ItemService) Get() string {
	res := cache.Read("K")

	if res == "" {
		data := cache.Data{K: "K", V: time.Now().GoString()}
		cache.Set(data)
		return "this key is empty"
	}

	return res
}
