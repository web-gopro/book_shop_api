package main

import (
	"context"
	"fmt"

	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/web-gopro/book_shop_api/api"
	"github.com/web-gopro/book_shop_api/config"
	"github.com/web-gopro/book_shop_api/pkg/db"
	"github.com/web-gopro/book_shop_api/redis"
	"github.com/web-gopro/book_shop_api/service"
)

func main() {
	cfg := config.Load()

	log := logger.NewLogger("", logger.LevelDebug)
	service := service.Service()

	fmt.Println(service)

	redisCli, err := db.ConnRedis(log, context.Background(), cfg.RedisConfig)

	if err != nil {

		return
	}

	fmt.Println(redisCli)

	cache := redis.NewRedisRepo(redisCli, log)

	engine := api.Api(api.Options{Service: service, Log: log, Cache: cache})

	engine.Run(":8080")
}
