package redis

import (
	"context"

	redis "github.com/go-redis/redis/v8"
	"github.com/humgal/art-server/util"
)

var ctx context.Context
var Rdb *redis.Client

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
	println(cmd.Result())
}
