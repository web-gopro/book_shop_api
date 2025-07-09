package handlers

import (
	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/web-gopro/book_shop_api/redis"
	"github.com/web-gopro/book_shop_api/service"
)

type Handler struct {
	service service.ServiceManagerI
	log     logger.LoggerI
	cache   redis.RedisRepoI
}

func NewHandlers(service service.ServiceManagerI, log logger.LoggerI, cache redis.RedisRepoI) Handler {

	return Handler{service: service, log: log, cache: cache}
}
