package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"profotdel-rest/pkg/response"
	"runtime/debug"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Warnf("PANIC recovered: %v\n%s", err, debug.Stack())

				msg := "panic occurred while processing the request"
				response.SendError(w, http.StatusInternalServerError, msg, response.ServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
