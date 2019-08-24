package handler

import (
	"github.com/micro/go-micro/client"
	"oauth_api/service"
	"sync"
	zap_log "utils/log"
)

var (
	log  *zap_log.Logger
	s    service.Service
	once sync.Once
)

func Init(cli client.Client) {
	once.Do(func() {
		log = zap_log.GetLoger()
		s = service.NewService(cli)
	})
}
