package jaeger

import (
	"github.com/opentracing/opentracing-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"io"
	"utils/config"
)

func NewTracer(svrName string) (opentracing.Tracer, io.Closer, error) {
	var cfg = &jaegerCfg.Configuration{}
	if err := config.C().App(cfg, "jaeger_tracer"); err != nil {
		panic(err)
	}
	cfg.ServiceName = svrName
	return cfg.NewTracer()
}
