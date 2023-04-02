package redis

import (
	"context"
	"fmt"

	"github.com/humgal/art-server/util"
	redis "github.com/redis/go-redis/v9"
	"github.com/rueian/rueidis"
)

var ctx context.Context
var Rdb *redis.Client

var RueidisClient *rueidis.Client

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

	//add for redis-search
	RueidisClient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	if err != nil {
		panic(err)
	}
	// SET key val NX
	RueidisClient.B().Ping().Build()
	pong, err := RueidisClient.Do(context.Background(), RueidisClient.B().Ping().Build()).ToString()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}
