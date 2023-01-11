package cacheOrd

import (
	"context"
	"gitoa.ru/go-4devs/cache"
	"gitoa.ru/go-4devs/cache/provider/memory"
	"nats"
)

type InMemory struct {
	mem *cache.Cache
}

func provider() cache.Provider {
	return memory.NewMap()
}

func (c *InMemory) InitCache() {
	c.mem = cache.New(provider())
}

func (c *InMemory) ReadOrder(uid string) (nats.Order, error) {
	ord := nats.Order{}
	ctx := context.Background()
	err := c.mem.Get(ctx, uid, &ord)
	return ord, err
}

func (c *InMemory) SaveOrder(ord *nats.Order) error {
	ctx := context.Background()
	return c.mem.Set(ctx, ord.OrderUid, *ord)
}

func (c *InMemory) RestoreCache(ords []nats.Order) error {
	for _, order := range ords {
		err := c.SaveOrder(&order)
		if err != nil {
			return err
		}
	}
	return nil
}
