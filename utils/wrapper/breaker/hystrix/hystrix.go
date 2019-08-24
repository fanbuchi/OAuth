package hystrix

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	zap_log "utils/log"

	status_code "utils/http"

	"github.com/afex/hystrix-go/hystrix"
)

var log *zap_log.Logger

func InitBreaker() {
	log = zap_log.GetLoger()
}

// BreakerWrapper hystrix breaker
func BreakerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI
		err := hystrix.Do(name, func() error {
			sct := &status_code.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
			h.ServeHTTP(sct.WrappedResponseWriter(), r)

			if sct.Status >= http.StatusBadRequest {
				str := fmt.Sprintf("status code %d", sct.Status)
				return errors.New(str)
			}
			return nil
		}, nil)
		if err != nil {
			log.Error("hystrix breaker err: ", zap.Error(err))
			return
		}
	})
}
