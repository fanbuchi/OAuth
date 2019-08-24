package main

import (
	"context"
	"github.com/micro/cli"
	"github.com/micro/go-config/encoder/yaml"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"go.uber.org/zap"
	"oauth_srv/handler"
	pbOauth "oauth_srv/proto/oauth"
	"oauth_srv/service"
	"os"
	"os/signal"
	"syscall"
	"utils/config"
	zap_log "utils/log"
)

var log *zap_log.Logger

func main() {
	// cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// New Service
	svr := grpc.NewService(
		micro.Name("zw.com.srv.oauth"),
		micro.Version("v1"),
		micro.Registry(etcdv3.NewRegistry(
			registry.Addrs("http://192.168.2.118:2379"),
		)),
		micro.Context(ctx),
		micro.Flags(
			cli.StringFlag{
				Name:  "cfg_path",
				Usage: "config path",
				Value: "oauth_svr.yaml",
			},
		),
		micro.Action(func(ctx *cli.Context) {
			// initialise config
			initCfg(ctx.String("cfg_path"))
			// initialise zap log
			zap_log.InitLogger()
			log = zap_log.GetLoger()
			// initialise service
			service.Init()
			// initialise handler
			handler.Init()

		}),
	)

	// Initialise service
	service.Init()

	// Register Handler
	if err := pbOauth.RegisterOauthSrvHandler(svr.Server(), new(handler.Oauth)); err != nil {
		log.Fatal("RegisterOauthSrvHandler", zap.Error(err))
	}

	//shutdown
	initShutdown(cancel)

	// Run service
	if err := svr.Run(); err != nil {
		if err == context.DeadlineExceeded {
			log.Info("service stopped")
		} else {
			log.Fatal("service run", zap.Error(err))
		}
	}
}

// initialise config
func initCfg(cfgAddr string) {
	enc := yaml.NewEncoder()
	cfgSource := file.NewSource(
		file.WithPath(cfgAddr),
		source.WithEncoder(enc),
	)
	config.Init(
		config.WithSource(cfgSource),
		config.WithApp("oauth_api"),
	)
	return
}

func initShutdown(cancel context.CancelFunc) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-c
		cancel()
		log.Info("shutting down")
	}()
}
