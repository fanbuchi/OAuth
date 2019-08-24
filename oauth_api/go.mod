module oauth_api

go 1.12

require (
	cloud.google.com/go v0.44.3 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/containerd/continuity v0.0.0-20190815185530-f2a389ac0a02 // indirect
	github.com/coreos/etcd v3.3.15+incompatible // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-micro v1.9.1
	github.com/micro/go-plugins v1.2.0
	github.com/micro/micro v1.9.1 // indirect
	github.com/miekg/dns v1.1.16 // indirect
	github.com/onsi/ginkgo v1.9.0 // indirect
	github.com/onsi/gomega v1.6.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	go.etcd.io/etcd v3.3.15+incompatible // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	golang.org/x/mobile v0.0.0-20190814143026-e8b3e6111d02 // indirect
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	golang.org/x/tools v0.0.0-20190822000311-fc82fb2afd64 // indirect
	google.golang.org/api v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.23.0 // indirect
	utils v0.0.0
)

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20190108154635-47c0da630f72
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
	utils => ../utils
)
