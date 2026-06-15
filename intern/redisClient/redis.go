package redisclient

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Rdb *redis.Client
	ctx context.Context
}

func NewRedisClient(addres string) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addres,
	})
	return &Client{
		Rdb: rdb,
		ctx: context.Background(),
	}
}
func (c *Client) Ping() error {
	ping, err := c.Rdb.Ping(c.ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println("redis is connected. ping:", ping)
	return nil
}
