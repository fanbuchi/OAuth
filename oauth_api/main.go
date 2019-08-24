package main

import (
	"context"
	go_hystrix "github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	"github.com/micro/go-config/encoder/yaml"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/file"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"go.uber.org/zap"
	"oauth_api/router"
	"os"
	"os/signal"
	"syscall"
	"utils/config"
	zap_log "utils/log"
	"utils/tracer/jaeger"
	tracer "utils/wrapper/tracer/opentracing/gin_micro"
)

var log *zap_log.Logger

func main() {
	// cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// create new web service
	service := web.NewService(
		web.Name("zw.com.web.oauth"),
		web.Registry(etcdv3.NewRegistry(
			registry.Addrs("http://192.168.2.118:2379"),
		)),
		web.Version("v1"),
		web.Context(ctx),
		web.Flags(
			cli.StringFlag{
				Name:  "cfg_path",
				Usage: "config path",
				Value: "oauth_api.yaml",
			},
		),
		web.Action(func(ctx *cli.Context) {
			//initialise config
			initCfg(ctx.String("cfg_path"))
			// initialise zap log
			zap_log.InitLogger()
			log = zap_log.GetLoger()

		}),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal("service init", zap.Error(err))
	}

	// initialise jaeger
	t, io, err := jaeger.NewTracer(service.Options().Name)
	if err != nil {
		log.Fatal("init jaeger tracer fault", zap.Error(err))
	}
	defer func() {
		err := io.Close()
		if err != nil {
			log.Error("jaeger io close", zap.Error(err))
		}
	}()
	tracer.InitTracer(t)

	// initialise client
	go_hystrix.DefaultTimeout = 5000
	sCli := service.Options().Service.Client()
	err = sCli.Init(
		client.Retries(3),
		client.Wrap(hystrix.NewClientWrapper()),
	)
	if err != nil {
		log.Fatal("service init fault", zap.Error(err))
	}

	// register html handler
	r := gin.Default()
	router.Init(r)
	service.Handle("/", r)

	//shutdown
	initShutdown(cancel)

	// run service
	if err := service.Run(); err != nil {
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
