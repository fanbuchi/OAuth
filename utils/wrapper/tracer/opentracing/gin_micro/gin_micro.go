package gin_micro

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	zap_log "utils/log"
)

const contextTracerKey = "Tracer-context"

var tracer opentracing.Tracer
var log *zap_log.Logger

// sf sampling frequency
var sf = 100

// SetSamplingFrequency 设置采样频率
// 0 <= n <= 100
func SetSamplingFrequency(n int) {
	sf = n
}

func InitTracer(t opentracing.Tracer) {
	tracer = t
	log = zap_log.GetLoger()
}

func TracerWrapper(c *gin.Context) {
	md := make(map[string]string)
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sp := tracer.StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
	defer sp.Finish()

	err := tracer.Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	if err != nil {
		log.Error("tracer inject ", zap.Error(err))
	}

	ctx := context.TODO()
	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	c.Set(contextTracerKey, ctx)

	c.Next()

	statusCode := c.Writer.Status()
	ext.HTTPStatusCode.Set(sp, uint16(statusCode))
	ext.HTTPMethod.Set(sp, c.Request.Method)
	ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
	if statusCode >= http.StatusInternalServerError {
		ext.Error.Set(sp, true)
	} else if rand.Intn(100) > sf {
		ext.SamplingPriority.Set(sp, 0)
	}
}

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {
	v, exist := c.Get(contextTracerKey)
	if exist == false {
		ok = false
		ctx = context.TODO()
		return
	}

	ctx, ok = v.(context.Context)
	return
}
