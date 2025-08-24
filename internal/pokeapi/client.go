package pokeapi

import (
	"time"
	"github.com/pjsmith404/gokedex/internal/pokecache"
)

func NewClient() Client {
	return Client{
		cache: pokecache.NewCache(5 * time.Minute),
	}
}
