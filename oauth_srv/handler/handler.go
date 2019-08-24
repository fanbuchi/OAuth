package handler

import (
	"oauth_srv/service"
	"sync"
	zap_log "utils/log"
)

var (
	log  *zap_log.Logger
	s    service.Service
	once sync.Once
)

func Init() {
	once.Do(func() {
		log = zap_log.GetLoger()
		s = service.NewService()
	})
}
