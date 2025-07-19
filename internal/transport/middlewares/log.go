package middlewares

import (
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"net/http"
	"time"
)

type responseRecover struct {
	http.ResponseWriter
	status int
}

func (r *responseRecover) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func LoggerMiddleware(log logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rr := &responseRecover{ResponseWriter: w, status: http.StatusOK}

			defer func() {
				log.Info(r.Context(), "request comleted",
					logger.Field{"request_time", start},
					logger.Field{"duration_ms", time.Since(start).Microseconds()},
					logger.Field{"request_method", r.Method},
					logger.Field{"request_path", r.URL.Path},
					logger.Field{"status", rr.status},
				)
			}()

			next.ServeHTTP(rr, r)
		})
	}
}
