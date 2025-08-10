package pokeapi

import (
	"time"
	"github.com/pjsmith404/gokedex/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
}

func NewClient() Client {
	return Client{
		cache: pokecache.NewCache(5 * time.Second),
	}
}
