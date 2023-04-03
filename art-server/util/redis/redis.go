package redis

import (
	"context"

	"github.com/humgal/art-server/util"
	redis "github.com/redis/go-redis/v9"
	"github.com/rueian/rueidis"
)

var ctx context.Context
var Rdb *redis.Client
var Client rueidis.Client

func init() {
	ctx = context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     util.Config.Redishost,
		Password: util.Config.Redispasswd,
		DB:       util.Config.Redisdb,
	})
	cmd := Rdb.Ping(ctx)
	if cmd.Err() != nil {
		panic(cmd.Err())
	}

	println(cmd.String())

	Client, _ = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	// SET key val NX
	pong, err := Client.Do(context.Background(), Client.B().Ping().Build()).ToString()
	if err != nil {
		panic(err)
	}
	println(pong)
}
