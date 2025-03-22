package routing

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"worker/internal/api/handling"
)

func SetUpRouting() (handler http.Handler) {
	mux := http.NewServeMux()

	mux.HandleFunc("/internal/api/worker/hash/crack/task", handling.HandleTaskRequest)

	handler = wrapWithLogging(mux)

	return
}

func Run(handler http.Handler) (err error) {
	log.Infof("Starting server at port 8081\n")

	err = http.ListenAndServe(":8081", handler)

	return
}

func wrapWithLogging(next http.Handler) (handler http.Handler) {
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("New request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})

	return
}
