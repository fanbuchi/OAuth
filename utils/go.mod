module utils

go 1.12

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-gonic/gin v1.4.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.9.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.15.0+incompatible
	go.uber.org/zap v1.9.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20190108154635-47c0da630f72
