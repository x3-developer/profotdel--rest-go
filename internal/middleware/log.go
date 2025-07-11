package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logrus.Infof("received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		logrus.Infof("response sent for: %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
