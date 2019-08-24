package service

import (
	"sync"
	zap_log "utils/log"
)

var (
	log  *zap_log.Logger
	once sync.Once
)

func Init() {
	once.Do(func() {
		log = zap_log.GetLoger()
	})
}

type Service interface {
}

type service struct {
}

func NewService() Service {
	return service{}
}
