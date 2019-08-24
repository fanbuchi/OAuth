package std_http

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
	status_code "utils/http"
	zap_log "utils/log"
)

// sf sampling frequency
var sf = 100

var tracer opentracing.Tracer
var log *zap_log.Logger

func init() {
	rand.Seed(time.Now().Unix())
}

func InitTracer(t opentracing.Tracer) {
	tracer = t
	log = zap_log.GetLoger()
	sf = 50
}

// SetSamplingFrequency 设置采样频率
// 0 <= n <= 100
func SetSamplingFrequency(n int) {
	sf = n
}

// TracerWrapper tracer wrapper
func TracerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		sp := tracer.StartSpan(r.URL.Path, opentracing.ChildOf(spanCtx))
		defer sp.Finish()

		err := tracer.Inject(sp.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Error("tracer inject", zap.Error(err))
		}

		sct := &status_code.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
		h.ServeHTTP(sct.WrappedResponseWriter(), r)

		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.EscapedPath())
		ext.HTTPStatusCode.Set(sp, uint16(sct.Status))
		if sct.Status >= http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		} else if rand.Intn(100) > sf {
			ext.SamplingPriority.Set(sp, 0)
		}
	})
}
